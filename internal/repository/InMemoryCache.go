package repository

import (
	"fmt"
	"sync"
)

type InMemoryCache struct {
	m     sync.RWMutex
	cache map[int]int
}

func (db *InMemoryCache) Set(key, value int) error {
	db.m.Lock()
	defer db.m.Unlock()
	db.cache[key] = value

	return nil
}

func (db *InMemoryCache) Get(key int) (int, error) {
	db.m.RLock()
	defer db.m.RUnlock()
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
