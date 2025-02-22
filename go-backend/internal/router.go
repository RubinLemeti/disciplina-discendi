package internal

import (
	"github.com/go-chi/chi/v5"
	"go-backend/internal/user"
)

func AggregatedRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Mount("/", user.UserRoutes())

	return router
}
