package api

import "net/http"

func (server *Server) cors(
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(
				"Access-Control-Allow-Origin",
				server.allowedOrigin,
			)
			w.Header().Set(
				"Access-Control-Allow-Headers",
				"Content-Type",
			)
			w.Header().Set(
				"Access-Control-Allow-Methods",
				"GET, POST, PATCH, DELETE, OPTIONS",
			)

			w.Header().Add("Vary", "Origin")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		},
	)
}
