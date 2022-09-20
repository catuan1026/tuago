package limits

import (
	"sync"
	"tuago"
)

type UniLimiter[K tuago.KeyAble] struct {
	mu   sync.Mutex
	dest map[K]bool
}

func NewUniLimiter[K tuago.KeyAble]() *UniLimiter[K] {
	return &UniLimiter[K]{
		dest: make(map[K]bool),
		mu:   sync.Mutex{},
	}
}

func (l *UniLimiter[K]) OK(key K) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	_, ok := l.dest[key]
	if !ok {
		l.dest[key] = true
	}
	return !ok
}

func (l *UniLimiter[K]) Release(key K) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if _, ok := l.dest[key]; ok {
		delete(l.dest, key)
	}
}
