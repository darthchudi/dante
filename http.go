package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type JSONErrorResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

func NewJSONErrorResponse(message string, err error) ([]byte, error) {
	errResponse := &JSONErrorResponse{Message: message, Error: err}
	return json.Marshal(errResponse)
}

func handleError(w http.ResponseWriter, httpCode int, err error, message string) {
	log.Printf("%v: %v\n", message, err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)

	errResponse, _ := NewJSONErrorResponse(message, err)
	_, _ = w.Write(errResponse)
}

func sendJSONResponse(w http.ResponseWriter, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to encode json response: %v", err))
	}
	return nil
}
