package types

import (
	"sync"
)

// OrderedMap represents a thread-safe ordered map.
type OrderedMap[K comparable, V any] struct {
	sync.RWMutex
	order []K
	items map[K]V
}

// NewOrderedMap creates a new instance of OrderedMap.
func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		order: make([]K, 0),
		items: make(map[K]V),
	}
}

// Put inserts a key-value pair into the map.
func (om *OrderedMap[K, V]) Put(key K, value V) {
	om.Lock()
	defer om.Unlock()

	if _, exists := om.items[key]; !exists {
		om.order = append(om.order, key)
	}

	om.items[key] = value
}

// Get retrieves the value associated with the given key.
func (om *OrderedMap[K, V]) Get(key K) (V, bool) {
	om.RLock()
	defer om.RUnlock()

	value, exists := om.items[key]
	return value, exists
}

// Remove deletes a key-value pair from the map.
func (om *OrderedMap[K, V]) Remove(key K) {
	om.Lock()
	defer om.Unlock()

	newOrder := []K{}
	for _, k := range om.order {
		if k != key {
			newOrder = append(newOrder, k)
		}
	}
	om.order = newOrder
	delete(om.items, key)
}

func (om *OrderedMap[K, V]) Length() int {
	om.RLock()
	defer om.RUnlock()
	return len(om.order)
}

// AsList returns the ordered list
func (om *OrderedMap[K, V]) Iterate(exec func(K, V)) {
	om.RLock()
	defer om.RUnlock()

	for _, v := range om.order {
		exec(v, om.items[v])
	}
}
