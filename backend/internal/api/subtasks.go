package api

import (
	"need-some-todo-overview/backend/internal/model"
	"net/http"
	"strings"
)

func (server *Server) createSubtask(
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

	subtask := model.Subtask{
		TodoID: todoID,
		Title:  title,
	}

	if err := server.repository.CreateSubtask(
		&subtask,
	); err != nil {
		serverError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, subtask)
}

func (server *Server) updateSubtask(
	w http.ResponseWriter,
	r *http.Request,
) {
	subtaskID, err := parsePathID(r, "subtaskID")
	if err != nil {
		badRequest(w, "Invalid subtask ID")
		return
	}

	if _, err := server.repository.FindSubtask(
		subtaskID,
	); err != nil {
		repositoryError(w, err, "Subtask not found")
		return
	}

	var input struct {
		Title *string `json:"title"`
		Done  *bool   `json:"done"`
	}

	if !decodeJSON(w, r, &input) {
		return
	}

	updates := make(map[string]any)

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

	subtask, err := server.repository.UpdateSubtask(
		subtaskID,
		updates,
	)
	if err != nil {
		repositoryError(w, err, "Subtask not found")
		return
	}

	writeJSON(w, http.StatusOK, subtask)
}

func (server *Server) deleteSubtask(
	w http.ResponseWriter,
	r *http.Request,
) {
	subtaskID, err := parsePathID(r, "subtaskID")
	if err != nil {
		badRequest(w, "Invalid subtask ID")
		return
	}

	if err := server.repository.DeleteSubtask(
		subtaskID,
	); err != nil {
		repositoryError(w, err, "Subtask not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
