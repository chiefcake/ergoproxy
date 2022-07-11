package storage

import "sync"

type Map[K comparable, V any] struct {
	m     map[K]V
	mutex *sync.Mutex
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		m:     make(map[K]V),
		mutex: &sync.Mutex{},
	}
}

func (m Map[K, V]) Insert(key K, value V) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.m[key] = value
}
