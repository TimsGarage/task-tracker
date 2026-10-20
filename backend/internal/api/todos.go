package api

import (
	"bytes"
	"encoding/json"
	"need-some-todo-overview/backend/internal/model"
	"net/http"
	"strings"
	"time"
)

func (server *Server) createTodo(
	w http.ResponseWriter,
	r *http.Request,
) {
	projectID, err := parsePathID(r, "projectID")
	if err != nil {
		badRequest(w, "Invalid project ID")
		return
	}

	if _, err := server.repository.FindProject(projectID); err != nil {
		repositoryError(w, err, "Project not found")
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

	todo := model.Todo{
		ProjectID: projectID,
		Title:     title,
		Priority:  model.PriorityMedium,
		Tags:      model.StringList{},
		Subtasks:  []model.Subtask{},
	}

	if err := server.repository.CreateTodo(&todo); err != nil {
		serverError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, todo)
}

func (server *Server) getTodo(
	w http.ResponseWriter,
	r *http.Request,
) {
	todoID, err := parsePathID(r, "todoID")
	if err != nil {
		badRequest(w, "Invalid todo ID")
		return
	}

	todo, err := server.repository.FindTodo(todoID)
	if err != nil {
		repositoryError(w, err, "Todo not found")
		return
	}

	writeJSON(w, http.StatusOK, todo)
}

func (server *Server) updateTodo(
	w http.ResponseWriter,
	r *http.Request,
) {
	todoID, err := parsePathID(r, "todoID")
	if err != nil {
		badRequest(w, "Invalid todo ID")
		return
	}

	if _, err := server.repository.FindTodo(todoID); err != nil {
		repositoryError(w, err, "Todo not found")
		return
	}

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

	updates := make(map[string]any)

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

		if _, err := server.repository.FindProject(
			*input.ProjectID,
		); err != nil {
			repositoryError(w, err, "Project not found")
			return
		}

		updates["project_id"] = *input.ProjectID
	}

	if input.Priority != nil {
		priority := strings.ToLower(
			strings.TrimSpace(*input.Priority),
		)

		if !model.ValidPriority(priority) {
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
		tags := model.StringList(
			normalizeTags(*input.Tags),
		)

		encodedTags, err := tags.JSON()
		if err != nil {
			serverError(w, err)
			return
		}

		// Explicit JSON text avoids passing a Go slice to SQLite.
		updates["tags"] = encodedTags
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

	todo, err := server.repository.UpdateTodo(todoID, updates)
	if err != nil {
		repositoryError(w, err, "Todo not found")
		return
	}

	writeJSON(w, http.StatusOK, todo)
}

func (server *Server) deleteTodo(
	w http.ResponseWriter,
	r *http.Request,
) {
	todoID, err := parsePathID(r, "todoID")
	if err != nil {
		badRequest(w, "Invalid todo ID")
		return
	}

	if err := server.repository.DeleteTodo(todoID); err != nil {
		repositoryError(w, err, "Todo not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func parseNullableDate(
	raw json.RawMessage,
) (*time.Time, error) {
	value := bytes.TrimSpace(raw)

	if bytes.Equal(value, []byte("null")) {
		return nil, nil
	}

	var dateText string
	if err := json.Unmarshal(value, &dateText); err != nil {
		return nil, err
	}

	dateText = strings.TrimSpace(dateText)
	if dateText == "" {
		return nil, nil
	}

	date, err := time.Parse(time.RFC3339, dateText)
	if err != nil {
		return nil, err
	}

	return &date, nil
}

func normalizeTags(tags []string) []string {
	result := make([]string, 0, len(tags))
	seen := make(map[string]bool)

	for _, value := range tags {
		tag := strings.TrimSpace(
			strings.TrimPrefix(value, "#"),
		)

		if tag == "" {
			continue
		}

		key := strings.ToLower(tag)
		if seen[key] {
			continue
		}

		seen[key] = true
		result = append(result, tag)
	}

	return result
}
