package api

import (
	"need-some-todo-overview/backend/internal/repository"
	"net/http"
)

type Server struct {
	repository    *repository.Repository
	allowedOrigin string
}

func New(
	repository *repository.Repository,
	allowedOrigin string,
) *Server {
	return &Server{
		repository:    repository,
		allowedOrigin: allowedOrigin,
	}
}

func (server *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc(
		"GET /api/projects",
		server.listProjects,
	)
	mux.HandleFunc(
		"POST /api/projects",
		server.createProject,
	)
	mux.HandleFunc(
		"PATCH /api/projects/{projectID}",
		server.updateProject,
	)
	mux.HandleFunc(
		"DELETE /api/projects/{projectID}",
		server.deleteProject,
	)

	mux.HandleFunc(
		"POST /api/projects/{projectID}/todos",
		server.createTodo,
	)

	mux.HandleFunc(
		"GET /api/todos/{todoID}",
		server.getTodo,
	)
	mux.HandleFunc(
		"PATCH /api/todos/{todoID}",
		server.updateTodo,
	)
	mux.HandleFunc(
		"DELETE /api/todos/{todoID}",
		server.deleteTodo,
	)

	mux.HandleFunc(
		"POST /api/todos/{todoID}/subtasks",
		server.createSubtask,
	)
	mux.HandleFunc(
		"PATCH /api/subtasks/{subtaskID}",
		server.updateSubtask,
	)
	mux.HandleFunc(
		"DELETE /api/subtasks/{subtaskID}",
		server.deleteSubtask,
	)

	mux.HandleFunc(
		"GET /api/health",
		server.health,
	)

	return server.cors(mux)
}

func (server *Server) health(
	w http.ResponseWriter,
	_ *http.Request,
) {
	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}
