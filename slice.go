package types

import (
	"sync"
)

type Slice[V any] struct {
	sync.RWMutex
	value []V
}

func (m *Slice[V]) Get() []V {
	m.RLock()
	defer m.RUnlock()
	return m.value
}

func (m *Slice[V]) Set(value []V) {
	m.Lock()
	m.value = value
	m.Unlock()
}

func (m *Slice[V]) Append(value ...V) {
	m.Lock()
	m.value = append(m.value, value...)
	m.Unlock()
}

func (m *Slice[V]) GetAndClear() []V {
	m.Lock()
	val := m.value
	m.value = []V{}
	m.Unlock()

	return val
}

func (m *Slice[V]) Pop() V {
	m.Lock()
	val := m.value[0]
	m.value = m.value[1:]
	m.Unlock()

	return val
}

func (m *Slice[V]) Length() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.value)
}

func NewSlice[V any]() *Slice[V] {
	return &Slice[V]{
		value: []V{},
	}
}
