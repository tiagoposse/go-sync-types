package types

import (
	"sync"
)

type Value[V any] struct {
	sync.RWMutex
	value V
}

func (m *Value[V]) Get() V {
	m.RLock()
	defer m.RUnlock()
	return m.value
}

func (m *Value[V]) Set(value V) {
	m.Lock()
	m.value = value
	m.Unlock()
}

func (m *Value[V]) GetAndClear() V {
	m.Lock()
	defer m.Unlock()
	val := m.value
	m.value = *new(V)
	return val
}
