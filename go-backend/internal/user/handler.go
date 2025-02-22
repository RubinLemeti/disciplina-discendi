package user

import (
	// "fmt"
	// "log/slog"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ErrorResponse struct {
	Key   string
	Error string
}

var validate *validator.Validate

func UserRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/users", GetUserList)
	router.Get("/users/{userId}", GetUserItem)
	router.Post("/users", AddUserItem)
	router.Patch("/users", UpdateUserItem)

	validate = validator.New(validator.WithRequiredStructEnabled())

	return router
}

func GetUserList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list of users"))
}

func GetUserItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user item"))
}

func AddUserItem(w http.ResponseWriter, r *http.Request) {
	var parsedRequestBody AddUserItemModel

	err := json.NewDecoder(r.Body).Decode(&parsedRequestBody)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = validate.Struct(parsedRequestBody)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(parsedRequestBody)
}

func UpdateUserItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user updated"))
}

func DeleteUserItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user deleted"))
}
