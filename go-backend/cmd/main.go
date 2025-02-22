package main

import (
	"go-backend/internal"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Mount("/", internal.AggregatedRoutes())

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("The service is running."))
	})

	slog.Info("Server is listening on port 8080")
	http.ListenAndServe(":8080", router)
}
