package api

import (
	"need-some-todo-overview/backend/internal/model"
	"net/http"
	"strings"
)

func (server *Server) listProjects(
	w http.ResponseWriter,
	_ *http.Request,
) {
	projects, err := server.repository.ListProjects()
	if err != nil {
		serverError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, projects)
}

func (server *Server) createProject(
	w http.ResponseWriter,
	r *http.Request,
) {
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

	project := model.Project{
		Title: title,
		Todos: []model.Todo{},
	}

	if err := server.repository.CreateProject(&project); err != nil {
		serverError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, project)
}

func (server *Server) updateProject(
	w http.ResponseWriter,
	r *http.Request,
) {
	projectID, err := parsePathID(r, "projectID")
	if err != nil {
		badRequest(w, "Invalid project ID")
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
		badRequest(w, "Project title is required")
		return
	}

	project, err := server.repository.UpdateProjectTitle(
		projectID,
		title,
	)
	if err != nil {
		repositoryError(w, err, "Project not found")
		return
	}

	writeJSON(w, http.StatusOK, project)
}

func (server *Server) deleteProject(
	w http.ResponseWriter,
	r *http.Request,
) {
	projectID, err := parsePathID(r, "projectID")
	if err != nil {
		badRequest(w, "Invalid project ID")
		return
	}

	if err := server.repository.DeleteProject(projectID); err != nil {
		repositoryError(w, err, "Project not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
