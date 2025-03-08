package service

import (
	"fmt"
	"net/http"
)

func (service *SimpleService) setValueHandler(w http.ResponseWriter, r *http.Request) {
	var key, value string

	if key = r.URL.Query().Get("key"); key == "" {
		http.Error(w, "BadRequest: expected 'key' in query parameters", http.StatusBadRequest)
		return
	}
	if value = r.URL.Query().Get("value"); value == "" {
		http.Error(w, "BadRequest: expected 'value' in query parameters", http.StatusBadRequest)
		return
	}

	service.repository.Set(key, value)
}

func (service *SimpleService) getValueHandler(w http.ResponseWriter, r *http.Request) {
	var key string

	if key = r.URL.Query().Get("key"); key == "" {
		http.Error(w, "BadRequest: expected 'key' in query parameters", http.StatusBadRequest)
		return
	}

	var value string
	var err error

	if value, err = service.repository.Get(key); err != nil {
		http.Error(w, fmt.Sprintf("key '%s' not found", key), http.StatusNotFound)
		return
	}

	w.Write([]byte(value))
}
