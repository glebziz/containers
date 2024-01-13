// Package omap implements an ordered map with a double linked list and a node pool.
//
// To iterate over a map (where m is a *OMap):
//
//	it := m.Iter()
//	for m.Next() {
//		// do something with it.Val()
//	}
package omap

import (
	"sync"

	"github.com/glebziz/containers/internal/node"
	"github.com/glebziz/containers/iter"
)

// OMap represents an ordered map.
// The zero value for OMap is an empty map ready to use.
type OMap[K comparable, V any] struct {
	data map[K]*node.Node[V]
	root node.Node[V]
	pool *node.Pool[V]

	m sync.RWMutex
}

// New returns an initialized map.
func New[K comparable, V any]() *OMap[K, V] {
	p := node.NewPool[V]()

	return &OMap[K, V]{
		data: make(map[K]*node.Node[V], p.Cap()),
		pool: p,
	}
}

// NewPresized returns an initialized map with an allocated pool of nodes.
func NewPresized[K comparable, V any](size int) *OMap[K, V] {
	return &OMap[K, V]{
		data: make(map[K]*node.Node[V], size),
		pool: node.NewPoolPresized[V](size),
	}
}

// Len returns the number of elements of map.
func (m *OMap[K, V]) Len() int {
	return len(m.data)
}

// Iter returns an iterator of the ordered map.
func (m *OMap[K, V]) Iter() *iter.Iter[V] {
	return iter.New[V](&m.root, iter.ForwardDir)
}

// Store stores the value by key in the map.
// The complexity is O(1).
func (m *OMap[K, V]) Store(key K, val V) {
	m.m.Lock()
	defer m.m.Unlock()

	if m.pool == nil {
		m.pool = node.NewPool[V]()
		m.data = make(map[K]*node.Node[V], m.pool.Cap())
	}

	if m.root.Next() == nil {
		m.root.SetNext(&m.root)
		m.root.SetPrev(&m.root)
	}

	n := m.pool.Pop()
	n.SetVal(val)
	m.root.Prev().Insert(n)

	m.data[key] = n
}

// Load returns the value by key from the map.
// The complexity is O(1).
func (m *OMap[K, V]) Load(key K) (val V, ok bool) {
	m.m.RLock()
	defer m.m.RUnlock()

	n, ok := m.data[key]
	return n.Val(), ok
}

// Delete removes the value by key from the map.
// The complexity is O(1).
func (m *OMap[K, V]) Delete(key K) {
	m.m.Lock()
	defer m.m.Unlock()

	n, ok := m.data[key]
	if ok {
		delete(m.data, key)
		n.Remove()
		m.pool.Push(n)
	}
}
