package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseRequestBody[T any](r *http.Request) (T, error) {
	var result T

	if r.Body == nil {
		return result, fmt.Errorf("request body is nil")
	}
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&result); err != nil {
		return result, fmt.Errorf("failed to decode request body: %w", err)
	}

	return result, nil
}

func HttpError(w http.ResponseWriter, err error, statusCode int) {
	fmt.Printf("HTTP Error: %v, Status Code: %d\n", err, statusCode)
	http.Error(w, err.Error(), statusCode)
}

func WriteResponseBody(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Printf("Failed to write response body: %v\n", err)
	}
}
