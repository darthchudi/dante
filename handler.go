package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gorilla/mux"
	"io"
	"log"
	"math/big"
	"net/http"
)

type GetItemResponse struct {
	Value string `json:"value"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error error `json:error`
}

func NewErrorResponse(message string, err error) ([]byte, error) {
	errResponse := &ErrorResponse{Message: message, Error: err}
	return json.Marshal(errResponse)
}

type SetItemRequestBody struct {
	Key uint `json:"key"`
	Value string `json:"value"`
}


func handleError(w http.ResponseWriter, httpCode int, err error, message string) {
	log.Printf("%v: %v\n", message, err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	errResponse, _ := NewErrorResponse(message, err)
	w.Write(errResponse)
}

func uintToBigInt(uint) *big.Int {
	return nil
}

func decodeRequest(body io.ReadCloser, v interface{}) error {
	err := json.NewDecoder(body).Decode(v)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to unmarshall json bytes: %v", err))
	}
	return nil
}

func encodeResponse(w http.ResponseWriter, v interface{}) error {
//	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to encode json response: %v", err))
	}
	return nil
}

func (s *Service) handleSetItem(w http.ResponseWriter, r *http.Request) {
	body := &SetItemRequestBody{}
	err := decodeRequest(r.Body, body); if err != nil {
		handleError(w, http.StatusInternalServerError, err, "failed to unmarshall json bytes")
		return
	}
	defer r.Body.Close()


	transaction, err := s.store.SetItem(&bind.TransactOpts{}, uintToBigInt(body.Key), body.Value)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err, "failed to set item")
		return
	}

	fmt.Println(transaction)
}

func (s *Service) handleGetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	if key == "1" {
		handleError(w, http.StatusBadRequest, errors.New("Key missing in request"), "Key missing in request")
		return
	}

	value, err := s.store.GetItem(&bind.CallOpts{}, uintToBigInt(0))
	if err != nil {
		handleError(w, http.StatusInternalServerError, err, fmt.Sprintf("failed to get item with key: %v", key))
		return
	}

	response := &GetItemResponse{Value: value}
	encodeResponse(w, response)
	return
}
