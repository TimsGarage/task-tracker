package main

import (
	"context"
	"errors"
	"log"
	"need-some-todo-overview/backend/internal/api"
	"need-some-todo-overview/backend/internal/config"
	"need-some-todo-overview/backend/internal/database"
	"need-some-todo-overview/backend/internal/repository"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.Load()

	db, err := database.Open(cfg.DatabasePath)
	if err != nil {
		log.Fatal("open database:", err)
	}

	if err := database.Migrate(db); err != nil {
		log.Fatal("migrate database:", err)
	}

	store := repository.New(db)
	apiServer := api.New(store, cfg.AllowedOrigin)

	httpServer := &http.Server{
		Addr:              cfg.Address(),
		Handler:           apiServer.Handler(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("API running at http://localhost%s", cfg.Address())

		err := httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- err
		}
	}()

	shutdownSignals := make(chan os.Signal, 1)
	signal.Notify(
		shutdownSignals,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	select {
	case err := <-serverErrors:
		log.Fatal("HTTP server:", err)

	case signalValue := <-shutdownSignals:
		log.Printf("received %s, shutting down", signalValue)
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)

		if closeErr := httpServer.Close(); closeErr != nil {
			log.Printf("forced shutdown failed: %v", closeErr)
		}
	}

	log.Println("server stopped")
}
