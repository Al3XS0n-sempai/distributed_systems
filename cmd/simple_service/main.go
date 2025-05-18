package main

import (
	"fmt"

	"github.com/Al3XS0n-sempai/distributed_systems/internal/repository"
	"github.com/Al3XS0n-sempai/distributed_systems/internal/service"
)

func main() {
	repo := repository.NewSyncMapInMemoryCache()
	service := service.NewSimpleService(repo)

	service.Init()

	if err := service.Run("0.0.0.0:8000"); err != nil {
		fmt.Printf("%v", err)
	}
}
