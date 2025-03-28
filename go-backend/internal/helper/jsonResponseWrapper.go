package helper

import (
	"encoding/json"
	"net/http"
	"time"
)

// Returns a generic json response with the given message and status code
// Any extra headers need to be added before calling the function
func GenericJsonResponse(w http.ResponseWriter, statusCode int, reponseMessage any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(reponseMessage)
}

type ResponseParamsObject[T any] struct {
	Data       *T
	Writer     http.ResponseWriter
	StatusCode int
	Path       string
	ErrorItem  ErrorSubModel
}

func ErrorJsonStandardResponse(params *ResponseParamsObject[any]) {
	if params.StatusCode == 0 {
		params.StatusCode = 500
	}

	if params.ErrorItem.Title == "" {
		params.ErrorItem.Title = "Internal Server Error"
	}

	if params.ErrorItem.Details == "" {
		params.ErrorItem.Details = "An unexpected error occurred"
	}

	params.Writer.Header().Set("Content-Type", "application/json")
	params.Writer.WriteHeader(params.StatusCode)

	response := FailureResponseModel{
		Error: ErrorSubModel{
			Title:   params.ErrorItem.Title,
			Details: params.ErrorItem.Details}, //err.Error()
		Path:       params.Path,
		Success:    false,
		StatusCode: params.StatusCode,
		Timestamp:  (time.Now()).Format("2006-01-02T03:04:05"),
	}

	json.NewEncoder(params.Writer).Encode(response)
}

func SuccessJsonStandardResponse(params *ResponseParamsObject[any]) {
	if params.StatusCode == 0 {
		params.StatusCode = 200
	}

	if params.ErrorItem.Title == "" {
		params.ErrorItem.Title = "Ok"
	}

	if params.ErrorItem.Details == "" {
		params.ErrorItem.Details = "Request completed successfully"
	}

	params.Writer.Header().Set("Content-Type", "application/json")
	params.Writer.WriteHeader(params.StatusCode)

	response := SuccessfulResponseModel{
		Path:       params.Path,
		Success:    true,
		StatusCode: params.StatusCode,
		Timestamp:  (time.Now()).Format("2006-01-02T03:04:05"),
	}

	json.NewEncoder(params.Writer).Encode(response)
}

func ItemJsonStandardResponse() {}

func ListJsonStandardResponse[T any](params *ResponseParamsObject[T]) {
	if params.StatusCode == 0 {
		params.StatusCode = 200
	}

	if params.ErrorItem.Title == "" {
		params.ErrorItem.Title = "Ok"
	}

	if params.ErrorItem.Details == "" {
		params.ErrorItem.Details = "Request completed successfully"
	}

	params.Writer.Header().Set("Content-Type", "application/json")
	params.Writer.WriteHeader(params.StatusCode)

	response := ListResponseModel[any]{
		Data:       params.Data,
		Path:       params.Path,
		Success:    true,
		StatusCode: params.StatusCode,
		Timestamp:  (time.Now()).Format("2006-01-02T03:04:05"),
	}

	json.NewEncoder(params.Writer).Encode(response)
}
