package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gorilla/mux"
	"math/big"
	"net/http"
	"strconv"
)

type SetItemRequest struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}

type GetItemResponse struct {
	Value string `json:"value"`
}

func (s *Service) handleSetItem(w http.ResponseWriter, r *http.Request) {
	body := &SetItemRequest{}
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err, "failed to unmarshall json bytes")
		return
	}
	defer r.Body.Close()

	transaction, err := s.store.SetItem(&bind.TransactOpts{}, big.NewInt(int64(body.Key)), body.Value)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err, "failed to set item")
		return
	}

	fmt.Printf("Transaction: %v", transaction)

	w.Write([]byte("Set item"))
}

func (s *Service) handleGetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["key"] == "" {
		handleError(w, http.StatusBadRequest, errors.New("key missing in request"), "Key missing in request")
		return
	}

	key, err := strconv.ParseInt(vars["key"], 10, 64)
	if err != nil {
		handleError(w, http.StatusBadRequest, err, "failed to parse key provided in request")
		return
	}

	value, err := s.store.GetItem(&bind.CallOpts{}, big.NewInt(key))
	if err != nil {
		handleError(w, http.StatusInternalServerError, err, fmt.Sprintf("failed to get item with key: %v", key))
		return
	}

	_ = sendJSONResponse(w, &GetItemResponse{Value: value})
	return
}
