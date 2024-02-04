package cache

import (
	"fmt"
	"sync"
)

type InMemoryCache struct {
	mu    *sync.RWMutex
	cache map[string]interface{}
}

func New() *InMemoryCache {
	return &InMemoryCache{cache: make(map[string]interface{}), mu: &sync.RWMutex{}}
}

func (r *InMemoryCache) Set(key string, value interface{}) error {
	_, ok := r.cache[key]
	if ok {
		return fmt.Errorf("value related with this <%s> key founded", key)
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	r.cache[key] = value

	return nil

}

func (r *InMemoryCache) Get(key string) (interface{}, error) {
	val, ok := r.cache[key]
	if !ok {
		return nil, fmt.Errorf("no value realted with this <%s> key", key)
	}

	return val, nil

}

func (r *InMemoryCache) Delete(key string) error {
	_, ok := r.cache[key]
	if !ok {
		return fmt.Errorf("no value realted with this <%s> key", key)
	}
	delete(r.cache, key)
	return nil

}
