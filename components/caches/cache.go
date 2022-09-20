package caches

import "tuago"

type CacheInf[K tuago.KeyAble, V any] interface {
	tuago.CleanAble
	Get(key K) (V, error)
	Del(key K) error
	Set(key K, value V) error
}
