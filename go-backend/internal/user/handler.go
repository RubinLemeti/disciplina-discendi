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
	router.Delete("/users/{userId}", userHandler.DeleteUserItem)

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

		if err = validate.Struct(GetUserListQueryParams{Limit: lStr}); err != nil {
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
			helper.ErrorJsonStandardResponseV2(w, http.StatusUnprocessableEntity, "/users", "Validation error", err.Error())
			return
		}

		err = validate.Struct(GetUserListQueryParams{Offset: oStr})
		if err != nil {
			slog.Error(err.Error())
			helper.ErrorJsonStandardResponseV2(w, http.StatusUnprocessableEntity, "/users", "Validation error", err.Error())
			return
		}

		offset = oStr
	}

	total, userList, err := h.UserService.GetUserList(limit, offset)

	if err != nil {
		slog.Error(err.Error())
		helper.ErrorJsonStandardResponseV2(w, 0, "/users", "", "")
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
		helper.ErrorJsonStandardResponseV2(w, http.StatusBadRequest, "/users/"+chi.URLParam(r, "userId"), "Validation error", "The id is not valid")
		return
	}

	userItemId := UserIdModel{Id: parsedUserId}

	// validate the entire struct
	err = validate.Struct(userItemId)
	if err != nil {
		slog.Error(err.Error())
		helper.ErrorJsonStandardResponseV2(w, http.StatusBadRequest, "/users/"+chi.URLParam(r, "userId"), "Validation error", "The id is not valid")
		return
	}

	// pass only the Id from the validated struct
	userItem, err := h.UserService.GetUserItem(userItemId.Id)
	if err != nil {
		slog.Error(err.Error())
		helper.ErrorJsonStandardResponseV2(w, 0, "/users/"+chi.URLParam(r, "userId"), "", "")
		return
	}
	if userItem == nil {
		slog.Error("User with id " + chi.URLParam(r, "userId") + "does not exist")
		helper.ErrorJsonStandardResponseV2(w, http.StatusNotFound, "/users/"+chi.URLParam(r, "userId"), "User not found", "The id is does not match with a user in the system")
		return
	}

	helper.ItemJsonStandardResponse(userItem, w, http.StatusOK, "/users/"+chi.URLParam(r, "userId"), "", "")
}

func (h UserHandler) AddUserItem(w http.ResponseWriter, r *http.Request) {
	var parsedRequestBody AddUserItemModel

	if err := json.NewDecoder(r.Body).Decode(&parsedRequestBody); err != nil {
		slog.Error(err.Error())
		helper.ErrorJsonStandardResponseV2(w, http.StatusNotFound, "/users", "", "")
		return
	}

	if err := validate.Struct(parsedRequestBody); err != nil {
		slog.Error(err.Error())
		helper.ErrorJsonStandardResponseV2(w, http.StatusUnprocessableEntity, "/users", "Validation error", err.Error())
		return
	}

	userId, err := h.UserService.AddUserItem(parsedRequestBody)
	if err != nil {

		if errors.Is(err, customerr.ErrUsernameNotUnique) {
			slog.Error(err.Error())
			helper.ErrorJsonStandardResponseV2(w, http.StatusUnprocessableEntity, "/users", "Username is taken", "Uniqueness constraint violation on the 'username' field.")
			return
		}

		slog.Error(err.Error())
		helper.ErrorJsonStandardResponseV2(w, 0, "/users", "", "")
		return
	}

	w.Header().Set("Location", "/users/"+strconv.Itoa(*userId))
	helper.SuccessJsonStandardResponseV2(w, http.StatusCreated, "/users")
}

func (h UserHandler) UpdateUserItem(w http.ResponseWriter, r *http.Request) {
	// parse the id url parameter
	parsedUserId, err := strconv.Atoi(chi.URLParam(r, "userId"))
	if err != nil {
		slog.Error(err.Error())
		helper.ErrorJsonStandardResponseV2(w, http.StatusBadRequest, "/users/"+chi.URLParam(r, "userId"), "Validation error", "The id is not a valid id")
		return
	}

	// validate the id url parameter
	if err = validate.Struct(UserIdModel{Id: parsedUserId}); err != nil {
		slog.Error(err.Error())
		helper.ErrorJsonStandardResponseV2(w, http.StatusBadRequest, "/users/"+chi.URLParam(r, "userId"), "Validation error", "The id is not a valid id")
		return
	}

	// validate that the json body is not empty
	if err = customval.ValidateNonEmptyRequest(r); err != nil {
		slog.Error(err.Error())
		helper.ErrorJsonStandardResponseV2(w, http.StatusBadRequest, "/users/"+chi.URLParam(r, "userId"), "Bad request", "The request is not valid json")
		return
	}

	// validate the json body
	var parsedRequestBody UpdateUserItemModel
	if err = json.NewDecoder(r.Body).Decode(&parsedRequestBody); err != nil {
		slog.Error(err.Error())
		helper.ErrorJsonStandardResponseV2(w, http.StatusBadRequest, "/users/"+chi.URLParam(r, "userId"), "Bad request", "The request is not valid json")
		return
	}

	if err = validate.Struct(parsedRequestBody); err != nil {
		slog.Error(err.Error())
		helper.ErrorJsonStandardResponseV2(w, http.StatusBadRequest, "/users/"+chi.URLParam(r, "userId"), "Validation error", err.Error())
		return
	}

	// update the variable
	userId, err := h.UserService.UpdateUserItem(parsedUserId, parsedRequestBody)
	if err != nil {

		if errors.Is(err, customerr.ErrUsernameNotUnique) {
			slog.Error(err.Error())
			helper.ErrorJsonStandardResponseV2(w, http.StatusUnprocessableEntity, "/users", "Username is taken", "Uniqueness constraint violation on the 'username' field.")
			return
		}

		slog.Error(err.Error())
		helper.ErrorJsonStandardResponseV2(w, 0, "/users", "", "")
		return
	}

	w.Header().Set("Location", "/users/"+strconv.Itoa(*userId))
	helper.SuccessJsonStandardResponseV2(w, 0, "/users/"+chi.URLParam(r, "userId"))
}

func (h UserHandler) DeleteUserItem(w http.ResponseWriter, r *http.Request) {
	// parse the id url parameter
	parsedUserId, err := strconv.Atoi(chi.URLParam(r, "userId"))
	if err != nil {
		slog.Error(err.Error())
		helper.ErrorJsonStandardResponseV2(w, http.StatusBadRequest, "/users/"+chi.URLParam(r, "userId"), "Validation error", "The id is not a valid id")
		return
	}

	// validate the id url parameter
	if err = validate.Struct(UserIdModel{Id: parsedUserId}); err != nil {
		slog.Error(err.Error())
		helper.ErrorJsonStandardResponseV2(w, http.StatusBadRequest, "/users/"+chi.URLParam(r, "userId"), "Validation error", "The id is not a valid id")
		return
	}

	userId, err := h.UserService.DeleteUserItem(parsedUserId)

	if userId == nil {
		helper.ErrorJsonStandardResponseV2(w, http.StatusNotFound, "/users/"+chi.URLParam(r, "userId"), "Validation error", "The id is not valid")
		helper.GenericJsonResponse(w, http.StatusNotFound, "User not found")
		return
	}

	if err != nil {
		slog.Error(err.Error())
		helper.ErrorJsonStandardResponseV2(w, 0, "/users/"+chi.URLParam(r, "userId"), "", "")
		return
	}

	w.Header().Set("Location", "/users/"+chi.URLParam(r, "userId"))
	helper.SuccessJsonStandardResponseV2(w, http.StatusNoContent, "/users/"+chi.URLParam(r, "userId"))
}
