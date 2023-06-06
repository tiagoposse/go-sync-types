package types

import (
	"sync"
)

// Map represents a thread-safe ordered map.
type Map[K comparable, V any] struct {
	sync.RWMutex
	items map[K]V
}

// NewMap creates a new instance of Map.
func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		items: make(map[K]V),
	}
}

// Put inserts a key-value pair into the map.
func (om *Map[K, V]) Put(key K, value V) {
	om.Lock()
	defer om.Unlock()

	om.items[key] = value
}

// Get retrieves the value associated with the given key.
func (om *Map[K, V]) Get(key K) (V, bool) {
	om.RLock()
	defer om.RUnlock()

	value, exists := om.items[key]
	return value, exists
}

// Remove deletes a key-value pair from the map.
func (om *Map[K, V]) Remove(key K) {
	om.Lock()
	defer om.Unlock()

	delete(om.items, key)
}

func (om *Map[K, V]) Length() int {
	om.RLock()
	defer om.RUnlock()

	return len(om.items)
}

// Iterate executes a function for each key, value pair
func (om *Map[K, V]) Iterate(exec func(K, V)) {
	om.RLock()
	defer om.RUnlock()

	for key, val := range om.items {
		exec(key, val)
	}
}
