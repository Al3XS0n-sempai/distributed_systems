package repository

import (
	"fmt"
	"sync"
)

type InMemoryCache struct {
	m     sync.RWMutex
	cache map[string]string
}

func (db *InMemoryCache) Set(key, value string) error {
	db.m.Lock()
	defer db.m.Unlock()
	db.cache[key] = value

	return nil
}

func (db *InMemoryCache) Get(key string) (string, error) {
	db.m.RLock()
	defer db.m.RUnlock()
	value, ok := db.cache[key]

	if !ok {
		return value, fmt.Errorf("no suck value with suck key (%s)", key)
	}

	return value, nil
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		cache: make(map[string]string, 100),
	}
}
