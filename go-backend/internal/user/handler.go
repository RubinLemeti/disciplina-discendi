package user

import (
	// "fmt"
	// "log/slog"
	"encoding/json"
	"errors"
	"go-backend/internal/helper"
	"go-backend/internal/helper/customerr"
	"go-backend/internal/helper/customval"
	"log/slog"
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
	router.Patch("/users/{userId}", userHandler.UpdateUserItem)

	validate = validator.New(validator.WithRequiredStructEnabled())

	return router
}

func (h UserHandler) GetUserList(w http.ResponseWriter, r *http.Request) {
	limit := 10
	offset := 0

	if l := r.URL.Query().Get("limit"); l != "" {
		lStr, err := strconv.Atoi(l)
		if err != nil {
			slog.Error(err.Error())
			helper.ErrorJsonStandardResponseV2(w, http.StatusUnprocessableEntity, "/users", "", "")
			return
		}

		
		if err = validate.Struct(GetUserListQueryParams{Limit: lStr});err != nil {
			slog.Error(err.Error())
			helper.ErrorJsonStandardResponseV2(w, http.StatusUnprocessableEntity, "/users", "Validation error", err.Error())
			return
		}

		limit = lStr
	}

	if o := r.URL.Query().Get("offset"); o != "" {
		oStr, err := strconv.Atoi(o)
		if err != nil {
			slog.Error(err.Error())
			helper.ErrorJsonStandardResponse(&helper.ResponseParamsObject[any]{
				StatusCode: http.StatusUnprocessableEntity,
				Writer:     w,
				Path:       "/users"})
			return
		}

		err = validate.Struct(GetUserListQueryParams{Offset: oStr})
		if err != nil {
			slog.Error(err.Error())
			helper.ErrorJsonStandardResponse(&helper.ResponseParamsObject[any]{
				StatusCode: http.StatusUnprocessableEntity,
				ErrorItem:  helper.ErrorSubModel{Title: "Validation error", Details: err.Error()},
				Writer:     w,
				Path:       "/users"})
			return
		}

		offset = oStr
	}

	total, userList, err := h.UserService.GetUserList(limit, offset)

	if err != nil {
		slog.Error(err.Error())
		helper.ErrorJsonStandardResponse(&helper.ResponseParamsObject[any]{
			Writer: w,
			Path:   "/users"})
		return
	}

	paginationMeta := helper.Paginate(*total, limit, offset, len(userList), "/users")

	helper.ListJsonStandardResponse(&helper.ResponseParamsObject[[]*User]{
		Data:   &userList,
		Meta:   *paginationMeta,
		Writer: w,
		Path:   "/users",
	})
}

func (h UserHandler) GetUserItem(w http.ResponseWriter, r *http.Request) {
	// Get URL parameter: userId
	// turn it from string to int
	parsedUserId, err := strconv.Atoi(chi.URLParam(r, "userId"))
	if err != nil {
		slog.Error(err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("The id is not a valid id")
		return
	}

	userItemId := UserIdModel{Id: parsedUserId}

	// validate the entire struct
	err = validate.Struct(userItemId)
	if err != nil {
		slog.Error(err.Error())
		helper.GenericJsonResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// pass only the Id from the validated struct
	userItem, err := h.UserService.GetUserItem(userItemId.Id)
	if err != nil {
		slog.Error(err.Error())
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
		slog.Error(err.Error())
		helper.ErrorJsonStandardResponse(&helper.ResponseParamsObject[any]{
			Writer:     w,
			StatusCode: http.StatusBadRequest,
			Path:       "/users"})
		return
	}

	err = validate.Struct(parsedRequestBody)
	if err != nil {
		slog.Error(err.Error())
		helper.ErrorJsonStandardResponse(&helper.ResponseParamsObject[any]{
			Writer:     w,
			StatusCode: http.StatusUnprocessableEntity,
			Path:       "/users",
			ErrorItem: helper.ErrorSubModel{
				Title:   "Validation error",
				Details: err.Error()}})
		return
	}

	userId, err := h.UserService.AddUserItem(parsedRequestBody)

	if err != nil {

		if errors.Is(err, customerr.ErrUsernameNotUnique) {
			slog.Error(err.Error())
			helper.ErrorJsonStandardResponse(&helper.ResponseParamsObject[any]{
				Writer:     w,
				StatusCode: http.StatusUnprocessableEntity,
				Path:       "/users",
				ErrorItem: helper.ErrorSubModel{
					Title:   "Username is taken",
					Details: "Uniqueness constraint violation on the 'username' field."}})
			return
		}

		slog.Error(err.Error())
		helper.ErrorJsonStandardResponse(&helper.ResponseParamsObject[any]{
			Writer: w,
			Path:   "/users",
		})
		return
	}

	w.Header().Set("Location", "/users/"+strconv.Itoa(*userId))
	helper.SuccessJsonStandardResponse(&helper.ResponseParamsObject[any]{
		Writer:     w,
		StatusCode: http.StatusCreated,
		Path:       "/users/" + strconv.Itoa(*userId),
	})
}

func (h UserHandler) UpdateUserItem(w http.ResponseWriter, r *http.Request) {
	parsedUserId, err := strconv.Atoi(chi.URLParam(r, "userId"))
	if err != nil {
		slog.Error(err.Error())
		helper.ErrorJsonStandardResponseV2(w, http.StatusBadRequest, "/users/{id}", "Validation error", "The id is not a valid id")
		return
	}

	err = validate.Struct(UserIdModel{Id: parsedUserId})
	if err != nil {
		slog.Error(err.Error())
		helper.ErrorJsonStandardResponse(&helper.ResponseParamsObject[any]{
			Writer:     w,
			StatusCode: http.StatusBadRequest,
			Path:       "/users"})
		return
	}

	var parsedRequestBody UpdateUserItemModel
	err = json.NewDecoder(r.Body).Decode(&parsedRequestBody)
	if err != nil {
		slog.Error(err.Error())
		helper.ErrorJsonStandardResponse(&helper.ResponseParamsObject[any]{
			Writer:     w,
			StatusCode: http.StatusBadRequest,
			Path:       "/users"})
		return
	}

	err = customval.ValidateNonEmptyRequest(r)
	if err != nil {
		slog.Error(err.Error())
		helper.ErrorJsonStandardResponse(&helper.ResponseParamsObject[any]{
			Writer:     w,
			StatusCode: http.StatusBadRequest,
			Path:       "/users"})
		return
	}

	err = validate.Struct(parsedRequestBody)
	if err != nil {
		slog.Error(err.Error())
		helper.ErrorJsonStandardResponse(&helper.ResponseParamsObject[any]{
			Writer:     w,
			StatusCode: http.StatusUnprocessableEntity,
			Path:       "/users",
			ErrorItem: helper.ErrorSubModel{
				Title:   "Validation error",
				Details: err.Error()}})
		return
	}

	w.Write([]byte("user updated"))
}

func (h UserHandler) DeleteUserItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user deleted"))
}
