package main

import (
	"go-backend/internal"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	// "log"
	"log/slog"
	"net/http"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("Failed to load ther environment variables")
		slog.Error(err.Error())
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Mount("/", internal.AggregatedRoutes())

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("The service is running."))
	})

	slog.Info("Server is listening on port 8080")
	http.ListenAndServe(":8080", router)
}
