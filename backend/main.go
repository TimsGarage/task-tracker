package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Project struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Todos []Todo `json:"todos" gorm:"constraint:OnDelete:CASCADE"`
}

type Todo struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	ProjectID uint   `json:"projectId" gorm:"not null;index"`
	Title     string `json:"title"`
	Done      bool   `json:"done"`
}

var db *gorm.DB

func main() {
	var err error

	db, err = gorm.Open(
		sqlite.Open("todos.db?_foreign_keys=on"),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&Project{}, &Todo{}); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/api/projects", projectsHandler)
	http.HandleFunc("/api/projects/", projectHandler)
	http.HandleFunc("/api/todos/", todoHandler)

	log.Println("API running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", cors(http.DefaultServeMux)))
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var projects []Project

		result := db.
			Order("id ASC").
			Preload("Todos", func(db *gorm.DB) *gorm.DB {
				return db.Order("id DESC")
			}).
			Find(&projects)

		if result.Error != nil {
			serverError(w, result.Error)
			return
		}

		writeJSON(w, http.StatusOK, projects)

	case http.MethodPost:
		var input struct {
			Title string `json:"title"`
		}

		if !decodeJSON(w, r, &input) {
			return
		}

		input.Title = strings.TrimSpace(input.Title)
		if input.Title == "" {
			http.Error(w, "Project title is required", http.StatusBadRequest)
			return
		}

		project := Project{
			Title: input.Title,
			Todos: []Todo{},
		}

		if err := db.Create(&project).Error; err != nil {
			serverError(w, err)
			return
		}

		writeJSON(w, http.StatusCreated, project)

	default:
		methodNotAllowed(w)
	}
}

func projectHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/projects/")
	parts := strings.Split(strings.Trim(path, "/"), "/")

	if len(parts) == 0 || parts[0] == "" {
		http.NotFound(w, r)
		return
	}

	projectID, err := parseID(parts[0])
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	// POST /api/projects/:id/todos
	if len(parts) == 2 && parts[1] == "todos" {
		createTodo(w, r, projectID)
		return
	}

	// PATCH or DELETE /api/projects/:id
	if len(parts) != 1 {
		http.NotFound(w, r)
		return
	}

	var project Project
	if err := db.First(&project, projectID).Error; err != nil {
		notFoundOrError(w, err, "Project not found")
		return
	}

	switch r.Method {
	case http.MethodPatch:
		var input struct {
			Title string `json:"title"`
		}

		if !decodeJSON(w, r, &input) {
			return
		}

		input.Title = strings.TrimSpace(input.Title)
		if input.Title == "" {
			http.Error(w, "Project title is required", http.StatusBadRequest)
			return
		}

		project.Title = input.Title

		if err := db.Save(&project).Error; err != nil {
			serverError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, project)

	case http.MethodDelete:
		if err := db.Delete(&project).Error; err != nil {
			serverError(w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		methodNotAllowed(w)
	}
}

func createTodo(w http.ResponseWriter, r *http.Request, projectID uint) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}

	var project Project
	if err := db.First(&project, projectID).Error; err != nil {
		notFoundOrError(w, err, "Project not found")
		return
	}

	var input struct {
		Title string `json:"title"`
	}

	if !decodeJSON(w, r, &input) {
		return
	}

	input.Title = strings.TrimSpace(input.Title)
	if input.Title == "" {
		http.Error(w, "Todo title is required", http.StatusBadRequest)
		return
	}

	todo := Todo{
		ProjectID: project.ID,
		Title:     input.Title,
	}

	if err := db.Create(&todo).Error; err != nil {
		serverError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, todo)
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	idText := strings.TrimPrefix(r.URL.Path, "/api/todos/")
	todoID, err := parseID(strings.Trim(idText, "/"))
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	var todo Todo
	if err := db.First(&todo, todoID).Error; err != nil {
		notFoundOrError(w, err, "Todo not found")
		return
	}

	switch r.Method {
	case http.MethodPatch:
		var input struct {
			Title *string `json:"title"`
			Done  *bool   `json:"done"`
		}

		if !decodeJSON(w, r, &input) {
			return
		}

		updates := map[string]any{}

		if input.Title != nil {
			title := strings.TrimSpace(*input.Title)
			if title == "" {
				http.Error(w, "Todo title is required", http.StatusBadRequest)
				return
			}

			updates["title"] = title
		}

		if input.Done != nil {
			updates["done"] = *input.Done
		}

		if len(updates) == 0 {
			http.Error(w, "Nothing to update", http.StatusBadRequest)
			return
		}

		if err := db.Model(&todo).Updates(updates).Error; err != nil {
			serverError(w, err)
			return
		}

		if err := db.First(&todo, todo.ID).Error; err != nil {
			serverError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, todo)

	case http.MethodDelete:
		if err := db.Delete(&todo).Error; err != nil {
			serverError(w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		methodNotAllowed(w)
	}
}

func parseID(value string) (uint, error) {
	id, err := strconv.ParseUint(value, 10, 64)
	return uint(id), err
}

func decodeJSON(w http.ResponseWriter, r *http.Request, value any) bool {
	if err := json.NewDecoder(r.Body).Decode(value); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return false
	}

	return true
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Println(err)
	}
}

func methodNotAllowed(w http.ResponseWriter) {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func serverError(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}

func notFoundOrError(w http.ResponseWriter, err error, message string) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, message, http.StatusNotFound)
		return
	}

	serverError(w, err)
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(
			"Access-Control-Allow-Origin",
			"http://localhost:5173",
		)
		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type",
		)
		w.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, PATCH, DELETE, OPTIONS",
		)

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
