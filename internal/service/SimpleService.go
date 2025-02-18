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
	service.mux.HandleFunc("POST /api/set/", service.setValueHandler)
	service.mux.HandleFunc("GET /api/get/{key}/", service.getValueHandler)
	// service.mux.HandleFunc("GET /api/swagger/", service.getSwagger)

	service.initialized = true
}

func (service *SimpleService) Run() error {
	if !service.initialized {
		return errors.New("service was not initialized before Run().\nCall Init() directly")
	}

	fmt.Println("Starting server at 0.0.0.0:8080")
	if err := http.ListenAndServe("0.0.0.0:8080", service.mux); err != nil {
		return err
	}

	return nil
}
