package internal

import (
	"github.com/go-chi/chi/v5"
	"go-backend/internal/user"
	"go-backend/internal/config"
)

func AggregatedRoutes() *chi.Mux {
	// Instantiate DB
	dbInstance := config.InitGORM()


	router := chi.NewRouter()
	router.Mount("/", user.UserRoutes(dbInstance))

	return router
}
