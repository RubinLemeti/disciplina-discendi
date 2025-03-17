package user

import (
	// "fmt"
	// "log/slog"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Key   string
	Error string
}

type UserHandler struct {
	UserService UserServiceI
}

func NewUserHandler(service UserServiceI) *UserHandler {
	return &UserHandler{UserService: service}
}

var validate *validator.Validate

func UserRoutes(db *gorm.DB) *chi.Mux {
	userRepository := NewUserRepository(db)
	userService := NewUserService(userRepository)
	userHandler := NewUserHandler(userService)

	router := chi.NewRouter()

	router.Get("/users", userHandler.GetUserList)
	router.Get("/users/{userId}", userHandler.GetUserItem)
	router.Post("/users", userHandler.AddUserItem)
	router.Patch("/users", userHandler.UpdateUserItem)

	validate = validator.New(validator.WithRequiredStructEnabled())

	return router
}

func (h UserHandler) GetUserList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list of users"))
}

func (h UserHandler) GetUserItem(w http.ResponseWriter, r *http.Request) {
	userItem, err := h.UserService.GetUserItem(1)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(userItem)
	// return
	// w.Write([]byte("user item"))
}

func (h UserHandler) AddUserItem(w http.ResponseWriter, r *http.Request) {
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

func (h UserHandler) UpdateUserItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user updated"))
}

func (h UserHandler) DeleteUserItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user deleted"))
}
