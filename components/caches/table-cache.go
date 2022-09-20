package caches

import (
	"errors"
	"sync"
	"tuago"
)

var (
	ErrCacheNotFound = errors.New("cache not found")
)

type TableCacheHandlerFunc[K tuago.KeyAble, V any] func(key K) (V, error)

func NewTableCache[K tuago.KeyAble, V any](handler TableCacheHandlerFunc[K, V]) *TableCache[K, V] {
	return &TableCache[K, V]{
		data:    make(map[K]V),
		handler: handler,
	}
}

type TableCache[K tuago.KeyAble, V any] struct {
	data    map[K]V
	mu      sync.RWMutex
	handler TableCacheHandlerFunc[K, V]
}

func (t *TableCache[K, V]) Get(key K) (V, error) {
	t.mu.RLock()
	val, ok := t.data[key]
	t.mu.RUnlock()
	if ok {
		return val, nil
	}
	if t.handler != nil {
		val, err := t.handler(key)
		if err != nil {
			return nil, err
		}
		err = t.Set(key, val)
		if err != nil {
			return nil, err
		}
		t.data[key] = val
		return val, nil
	}
	return val, ErrCacheNotFound
}

func (t *TableCache[K, V]) Set(key K, value V) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.data[key] = value
	return nil
}

func (t *TableCache[K, V]) Del(key K) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	if _, ok := t.data[key]; ok {
		delete(t.data, key)
	}
	return nil
}

func (t *TableCache[K, V]) Clean() error {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.data = make(map[K]V)
	return nil
}
