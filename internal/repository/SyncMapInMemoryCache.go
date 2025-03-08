package repository

import (
	"fmt"
	"sync"
)

type SyncMapInMemoryCache struct {
	cache sync.Map
}

func (db *SyncMapInMemoryCache) Set(key, value string) error {
	db.cache.Store(key, value)

	return nil
}

func (db *SyncMapInMemoryCache) Get(key string) (string, error) {
	value, ok := db.cache.Load(key)

	if !ok {
		return "", fmt.Errorf("no suck value with suck key (%s)", key)
	}

	return value.(string), nil
}

func NewSyncMapInMemoryCache() *SyncMapInMemoryCache {
	return &SyncMapInMemoryCache{
		cache: sync.Map{},
	}
}
