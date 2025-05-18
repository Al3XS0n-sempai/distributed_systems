package main

import (
	"fmt"

	"github.com/Al3XS0n-sempai/distributed_systems/internal/repository"
	"github.com/Al3XS0n-sempai/distributed_systems/internal/service"
)

func main() {
	// repo := repository.NewSyncMapInMemoryCache()
	connStrings := []string{
		"postgres://user:pass@localhost:5432/shard1?sslmode=disable",
		"postgres://user:pass@localhost:5433/shard2?sslmode=disable",
		"postgres://user:pass@localhost:5434/shard3?sslmode=disable",
		"postgres://user:pass@localhost:5435/shard4?sslmode=disable",
		"postgres://user:pass@localhost:5436/shard5?sslmode=disable",
		// "postgres://user:pass@localhost:5437/shard6?sslmode=disable",
	}
	repo, err := repository.NewShardedPostgresql(connStrings)
	if err != nil {
		fmt.Println("Can't create repository:\n", err)
		return
	}
	service := service.NewSimpleService(repo)

	// CHECK THAT KEYS ARE DIFFERENT
	// repo.Get("1")
	// repo.Get("123")
	// repo.Get("3214")
	// repo.Get("81283")
	// repo.Get("57812")
	// repo.Get("12738")
	// repo.Get("4712")

	service.Init()

	if err := service.Run("0.0.0.0:8000"); err != nil {
		fmt.Printf("%v", err)
	}
}
