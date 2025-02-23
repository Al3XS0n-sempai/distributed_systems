package service

import (
	"errors"
	"fmt"
	"net/http"
)

type SimpleServiceRepository interface {
	Set(key, value int) error
	Get(key int) (int, error)
}

// SimpleService is a simple service that listens 2 HTTP endpoints:
//   - /api/set
//   - /api/get
//   - /api/swagger
//
// API Specification can be checked in swagger file in /api/REST directory
type SimpleService struct {
	mux        *http.ServeMux
	repository SimpleServiceRepository

	initialized bool
}

func NewSimpleService(repo SimpleServiceRepository) *SimpleService {
	if repo == nil {
		panic("repo can't be nil")
	}

	return &SimpleService{
		mux:         http.NewServeMux(),
		repository:  repo,
		initialized: false,
	}
}

func (service *SimpleService) Init() {
	service.mux.HandleFunc("POST /put", service.setValueHandler)
	service.mux.HandleFunc("GET /get", service.getValueHandler)
	// service.mux.HandleFunc("GET /api/swagger/", service.getSwagger)

	service.initialized = true
}

func (service *SimpleService) Run(addr string) error {
	if !service.initialized {
		return errors.New("service was not initialized before Run().\nCall Init() directly")
	}

	fmt.Printf("Starting server at %s", addr)
	if err := http.ListenAndServe(addr, service.mux); err != nil {
		return err
	}

	return nil
}
