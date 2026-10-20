package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Project struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Todos []Todo `json:"todos" gorm:"constraint:OnDelete:CASCADE"`
}

type Todo struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	ProjectID uint       `json:"projectId" gorm:"not null;index"`
	Title     string     `json:"title"`
	Done      bool       `json:"done"`
	DueDate   *time.Time `json:"dueDate"`
	Priority  string     `json:"priority" gorm:"default:medium"`
	Notes     string     `json:"notes"`
	Tags      []string   `json:"tags" gorm:"serializer:json"`
	Subtasks  []Subtask  `json:"subtasks" gorm:"constraint:OnDelete:CASCADE"`
}

type Subtask struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	TodoID uint   `json:"todoId" gorm:"not null;index"`
	Title  string `json:"title"`
	Done   bool   `json:"done"`
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

	if err := db.AutoMigrate(
		&Project{},
		&Todo{},
		&Subtask{},
	); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/api/projects", projectsHandler)
	http.HandleFunc("/api/projects/", projectHandler)
	http.HandleFunc("/api/todos/", todoHandler)
	http.HandleFunc("/api/subtasks/", subtaskHandler)

	log.Println("API running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(
		":8080",
		cors(http.DefaultServeMux),
	))
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var projects []Project

		err := db.
			Order("id ASC").
			Preload("Todos", func(query *gorm.DB) *gorm.DB {
				return query.Order("id DESC")
			}).
			Preload("Todos.Subtasks", func(query *gorm.DB) *gorm.DB {
				return query.Order("id ASC")
			}).
			Find(&projects).
			Error

		if err != nil {
			serverError(w, err)
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

		title := strings.TrimSpace(input.Title)
		if title == "" {
			badRequest(w, "Project title is required")
			return
		}

		project := Project{
			Title: title,
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
	path := strings.Trim(
		strings.TrimPrefix(r.URL.Path, "/api/projects/"),
		"/",
	)
	parts := strings.Split(path, "/")

	if len(parts) == 0 || parts[0] == "" {
		http.NotFound(w, r)
		return
	}

	projectID, err := parseID(parts[0])
	if err != nil {
		badRequest(w, "Invalid project ID")
		return
	}

	if len(parts) == 2 && parts[1] == "todos" {
		createTodo(w, r, projectID)
		return
	}

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

		title := strings.TrimSpace(input.Title)
		if title == "" {
			badRequest(w, "Project title is required")
			return
		}

		if err := db.Model(&project).
			Update("title", title).
			Error; err != nil {
			serverError(w, err)
			return
		}

		project.Title = title
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

func createTodo(
	w http.ResponseWriter,
	r *http.Request,
	projectID uint,
) {
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

	title := strings.TrimSpace(input.Title)
	if title == "" {
		badRequest(w, "Todo title is required")
		return
	}

	todo := Todo{
		ProjectID: project.ID,
		Title:     title,
		Priority:  "medium",
		Tags:      []string{},
		Subtasks:  []Subtask{},
	}

	if err := db.Create(&todo).Error; err != nil {
		serverError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, todo)
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(
		strings.TrimPrefix(r.URL.Path, "/api/todos/"),
		"/",
	)
	parts := strings.Split(path, "/")

	if len(parts) == 0 || parts[0] == "" {
		http.NotFound(w, r)
		return
	}

	todoID, err := parseID(parts[0])
	if err != nil {
		badRequest(w, "Invalid todo ID")
		return
	}

	if len(parts) == 2 && parts[1] == "subtasks" {
		createSubtask(w, r, todoID)
		return
	}

	if len(parts) != 1 {
		http.NotFound(w, r)
		return
	}

	var todo Todo
	if err := loadTodo(todoID, &todo); err != nil {
		notFoundOrError(w, err, "Todo not found")
		return
	}

	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, todo)

	case http.MethodPatch:
		updateTodo(w, r, &todo)

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

func updateTodo(
	w http.ResponseWriter,
	r *http.Request,
	todo *Todo,
) {
	var input struct {
		Title     *string         `json:"title"`
		Done      *bool           `json:"done"`
		ProjectID *uint           `json:"projectId"`
		DueDate   json.RawMessage `json:"dueDate"`
		Priority  *string         `json:"priority"`
		Notes     *string         `json:"notes"`
		Tags      *[]string       `json:"tags"`
	}

	if !decodeJSON(w, r, &input) {
		return
	}

	updates := map[string]any{}

	if input.Title != nil {
		title := strings.TrimSpace(*input.Title)

		if title == "" {
			badRequest(w, "Todo title is required")
			return
		}

		updates["title"] = title
	}

	if input.Done != nil {
		updates["done"] = *input.Done
	}

	if input.ProjectID != nil {
		if *input.ProjectID == 0 {
			badRequest(w, "A project is required")
			return
		}

		var project Project
		if err := db.First(&project, *input.ProjectID).Error; err != nil {
			notFoundOrError(w, err, "Project not found")
			return
		}

		updates["project_id"] = *input.ProjectID
	}

	if input.Priority != nil {
		priority := strings.ToLower(
			strings.TrimSpace(*input.Priority),
		)

		if !validPriority(priority) {
			badRequest(
				w,
				"Priority must be low, medium, or high",
			)
			return
		}

		updates["priority"] = priority
	}

	if input.Notes != nil {
		updates["notes"] = strings.TrimSpace(*input.Notes)
	}

	if input.Tags != nil {
		updates["tags"] = normalizeTags(*input.Tags)
	}

	if len(input.DueDate) > 0 {
		dueDate, err := parseNullableDate(input.DueDate)
		if err != nil {
			badRequest(w, "Invalid due date")
			return
		}

		updates["due_date"] = dueDate
	}

	if len(updates) == 0 {
		badRequest(w, "Nothing to update")
		return
	}

	if err := db.Model(todo).Updates(updates).Error; err != nil {
		serverError(w, err)
		return
	}

	if err := loadTodo(todo.ID, todo); err != nil {
		serverError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, todo)
}

func createSubtask(
	w http.ResponseWriter,
	r *http.Request,
	todoID uint,
) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}

	var todo Todo
	if err := db.First(&todo, todoID).Error; err != nil {
		notFoundOrError(w, err, "Todo not found")
		return
	}

	var input struct {
		Title string `json:"title"`
	}

	if !decodeJSON(w, r, &input) {
		return
	}

	title := strings.TrimSpace(input.Title)
	if title == "" {
		badRequest(w, "Subtask title is required")
		return
	}

	subtask := Subtask{
		TodoID: todo.ID,
		Title:  title,
	}

	if err := db.Create(&subtask).Error; err != nil {
		serverError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, subtask)
}

func subtaskHandler(w http.ResponseWriter, r *http.Request) {
	idText := strings.Trim(
		strings.TrimPrefix(r.URL.Path, "/api/subtasks/"),
		"/",
	)

	subtaskID, err := parseID(idText)
	if err != nil {
		badRequest(w, "Invalid subtask ID")
		return
	}

	var subtask Subtask
	if err := db.First(&subtask, subtaskID).Error; err != nil {
		notFoundOrError(w, err, "Subtask not found")
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
				badRequest(w, "Subtask title is required")
				return
			}

			updates["title"] = title
		}

		if input.Done != nil {
			updates["done"] = *input.Done
		}

		if len(updates) == 0 {
			badRequest(w, "Nothing to update")
			return
		}

		if err := db.Model(&subtask).
			Updates(updates).
			Error; err != nil {
			serverError(w, err)
			return
		}

		if err := db.First(&subtask, subtask.ID).
			Error; err != nil {
			serverError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, subtask)

	case http.MethodDelete:
		if err := db.Delete(&subtask).Error; err != nil {
			serverError(w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		methodNotAllowed(w)
	}
}

func loadTodo(id uint, todo *Todo) error {
	return db.
		Preload("Subtasks", func(query *gorm.DB) *gorm.DB {
			return query.Order("id ASC")
		}).
		First(todo, id).
		Error
}

func parseNullableDate(
	raw json.RawMessage,
) (*time.Time, error) {
	trimmed := bytes.TrimSpace(raw)

	if bytes.Equal(trimmed, []byte("null")) {
		return nil, nil
	}

	var value string
	if err := json.Unmarshal(trimmed, &value); err != nil {
		return nil, err
	}

	value = strings.TrimSpace(value)
	if value == "" {
		return nil, nil
	}

	date, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return nil, err
	}

	return &date, nil
}

func validPriority(priority string) bool {
	return priority == "low" ||
		priority == "medium" ||
		priority == "high"
}

func normalizeTags(tags []string) []string {
	result := make([]string, 0, len(tags))
	seen := map[string]bool{}

	for _, value := range tags {
		tag := strings.TrimSpace(
			strings.TrimPrefix(value, "#"),
		)

		if tag == "" || seen[tag] {
			continue
		}

		seen[tag] = true
		result = append(result, tag)
	}

	return result
}

func parseID(value string) (uint, error) {
	id, err := strconv.ParseUint(value, 10, 64)
	return uint(id), err
}

func decodeJSON(
	w http.ResponseWriter,
	r *http.Request,
	value any,
) bool {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(value); err != nil {
		badRequest(w, "Invalid JSON")
		return false
	}

	return true
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if value != nil {
		if err := json.NewEncoder(w).Encode(value); err != nil {
			log.Println(err)
		}
	}
}

func badRequest(w http.ResponseWriter, message string) {
	http.Error(w, message, http.StatusBadRequest)
}

func methodNotAllowed(w http.ResponseWriter) {
	http.Error(
		w,
		"Method not allowed",
		http.StatusMethodNotAllowed,
	)
}

func serverError(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(
		w,
		"Internal server error",
		http.StatusInternalServerError,
	)
}

func notFoundOrError(
	w http.ResponseWriter,
	err error,
	message string,
) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, message, http.StatusNotFound)
		return
	}

	serverError(w, err)
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
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
		},
	)
}
