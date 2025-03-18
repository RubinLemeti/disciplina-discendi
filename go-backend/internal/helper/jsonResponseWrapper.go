package helper

import (
	"encoding/json"
	"net/http"
)

func GenericJsonResponse(w http.ResponseWriter, statusCode int, reponseMessage any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(reponseMessage)
}
