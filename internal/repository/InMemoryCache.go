package repository

import (
	"fmt"
)

type InMemoryCache struct {
	cache map[int]int
}

func (db *InMemoryCache) Set(key, value int) error {
	db.cache[key] = value

	return nil
}

func (db InMemoryCache) Get(key int) (int, error) {
	value, ok := db.cache[key]

	if !ok {
		return value, fmt.Errorf("no suck value with suck key (%d)", key)
	}

	return value, nil
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		cache: make(map[int]int, 100),
	}
}
