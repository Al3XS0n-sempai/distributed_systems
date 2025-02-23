package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (service *SimpleService) setValueHandler(w http.ResponseWriter, r *http.Request) {
	var sKey, sValue string

	if sKey = r.URL.Query().Get("key"); sKey == "" {
		http.Error(w, "BadRequest: expected 'key' in query parameters", http.StatusBadRequest)
		return
	}
	if sValue = r.URL.Query().Get("value"); sValue == "" {
		http.Error(w, "BadRequest: expected 'value' in query parameters", http.StatusBadRequest)
		return
	}

	var key, value int
	var err error

	if key, err = strconv.Atoi(sKey); err != nil {
		http.Error(w, "'key' should be a valid integer number", http.StatusBadRequest)
		return
	}

	if value, err = strconv.Atoi(sValue); err != nil {
		http.Error(w, "'value' should be a valid integer number", http.StatusBadRequest)
		return
	}

	service.repository.Set(key, value)
}

func (service *SimpleService) getValueHandler(w http.ResponseWriter, r *http.Request) {
	var sKey string

	if sKey = r.URL.Query().Get("key"); sKey == "" {
		http.Error(w, "BadRequest: expected 'key' in query parameters", http.StatusBadRequest)
		return
	}

	var key, value int
	var err error

	if key, err = strconv.Atoi(sKey); err != nil {
		http.Error(w, "'key' should be a valid integer number", http.StatusBadRequest)
		return
	}

	if value, err = service.repository.Get(key); err != nil {
		http.Error(w, fmt.Sprintf("key %d not found", key), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"value": value})
}
