package shared

import (
	"encoding/json"
	"net/http"
)

func WriteError(w http.ResponseWriter, status int, code string, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(APIError{
		Code:    code,
		Message: msg,
	})
}
