package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (service *SimpleService) setValueHandler(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Key   int `json:"key"`
		Value int `json:"value"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Wrong json, should be 2 keys of type int: key and value", http.StatusBadRequest)
		return
	}

	if err := service.repository.Set(payload.Key, payload.Value); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (service *SimpleService) getValueHandler(w http.ResponseWriter, r *http.Request) {
	var skey string

	if skey = r.PathValue("key"); skey == "" {
		http.Error(w, "BadRequest: expected path /api/get/{key}, where key is an interger number", http.StatusBadRequest)
		return
	}

	var key, value int
	var err error

	if key, err = strconv.Atoi(skey); err != nil {
		http.Error(w, "key in /api/get/{key}, should be a valid integer number", http.StatusBadRequest)
		return
	}

	if value, err = service.repository.Get(key); err != nil {
		http.Error(w, fmt.Sprintf("key %d not found", key), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"value": value})
	w.WriteHeader(http.StatusOK)
}
