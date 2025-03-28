package user

import (
	// "fmt"
	// "log/slog"
	"encoding/json"
	"errors"
	"go-backend/internal/helper"
	"go-backend/internal/helper/customerr"
	"net/http"
	"strconv"

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
	// Get URL parameter: userId
	// turn it from string to int
	parsedUserId, err := strconv.Atoi(chi.URLParam(r, "userId"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("The id is not a valid id")
		return
	}

	userItemId := UserIdModel{Id: parsedUserId}

	// validate the entire struct
	err = validate.Struct(userItemId)
	if err != nil {
		helper.GenericJsonResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// pass only the Id from the validated struct
	userItem, err := h.UserService.GetUserItem(userItemId.Id)
	if err != nil {
		helper.GenericJsonResponse(w, http.StatusBadRequest, "Issue with the query")
		return
	}
	if userItem == nil {
		helper.GenericJsonResponse(w, http.StatusNotFound, "User not found")
		return
	}

	helper.GenericJsonResponse(w, http.StatusOK, userItem)
}

func (h UserHandler) AddUserItem(w http.ResponseWriter, r *http.Request) {
	var parsedRequestBody AddUserItemModel

	err := json.NewDecoder(r.Body).Decode(&parsedRequestBody)

	if err != nil {
		helper.GenericJsonResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = validate.Struct(parsedRequestBody)

	if err != nil {
		helper.GenericJsonResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	userId, err := h.UserService.AddUserItem(parsedRequestBody)

	if err != nil {
		if errors.Is(err, customerr.ErrUsernameNotUnique) {

			helper.ErrorJsonStandardResponse(&helper.ResponseParamsObject{
				Writer:     w,
				StatusCode: http.StatusUnprocessableEntity,
				Path:       "/users",
				ErrorItem: helper.ErrorSubModel{
					Title:   "Username is taken",
					Details: "Uniqueness constraint violation on the 'username' field."}})
			return
		}

		helper.ErrorJsonStandardResponse(&helper.ResponseParamsObject{
			Writer: w,
			Path:   "/users",
		})
		return
	}

	w.Header().Set("Location", "/users/"+strconv.Itoa(*userId))
	helper.SuccessJsonStandardResponse(&helper.ResponseParamsObject{
		Writer:     w,
		StatusCode: http.StatusCreated,
		Path:       "/users/" + strconv.Itoa(*userId),
	})
}

func (h UserHandler) UpdateUserItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user updated"))
}

func (h UserHandler) DeleteUserItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user deleted"))
}
