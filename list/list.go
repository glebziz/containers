// Package list implements a doubly linked list with a pool of nodes.
//
// To iterate over a list (where l is a *List):
//
//	it := l.Iter()
//	for it.Next() {
//		// do something with it.Val()
//	}
package list

import (
	"sync"

	"github.com/glebziz/containers/internal/node"
	"github.com/glebziz/containers/iter"
)

// List represents a doubly linked list.
// The zero value for List is an empty list ready to use.
type List[T any] struct {
	root node.Node[T]
	pool *node.Pool[T]
	m    sync.RWMutex
	len  int
}

// New returns an initialized list.
func New[T any]() *List[T] {
	l := &List[T]{
		pool: node.NewPool[T](),
	}

	return l
}

// NewPresized returns an initialized list with an allocated pool of nodes.
func NewPresized[T any](size int) *List[T] {
	l := &List[T]{
		pool: node.NewPoolPresized[T](size),
	}

	return l
}

// Len returns the number of elements of list.
func (l *List[T]) Len() int {
	return l.len
}

// Iter returns a list iterator with forward direction.
func (l *List[T]) Iter() *iter.Iter[T] {
	return iter.New[T](&l.root, iter.ForwardDir)
}

// RIter returns a list iterator with reverse direction.
func (l *List[T]) RIter() *iter.Iter[T] {
	return iter.New[T](&l.root, iter.ReverseDir)
}

// Front returns the value of the first element of the list or zero value if the list is empty.
func (l *List[T]) Front() T {
	l.m.RLock()
	defer l.m.RUnlock()

	return l.root.Next().Val()
}

// Back returns the value of the last element of the list or zero value if the list is empty.
func (l *List[T]) Back() T {
	l.m.RLock()
	defer l.m.RUnlock()

	return l.root.Prev().Val()
}

// Get returns the value of the i-th element of the list or zero value if the list is empty or len < i.
// The complexity is O(n).
func (l *List[T]) Get(i int) T {
	l.m.RLock()
	defer l.m.RUnlock()

	return l.get(i).Val()
}

// PushFront inserts a new value at the front of the list.
func (l *List[T]) PushFront(v T) {
	l.m.Lock()
	defer l.m.Unlock()

	l.lazyInit()
	l.insert(v, &l.root)
}

// PushBack inserts a new value at the back of the list.
func (l *List[T]) PushBack(v T) {
	l.m.Lock()
	defer l.m.Unlock()

	l.lazyInit()
	l.insert(v, l.root.Prev())
}

// PushAfter inserts a new value after the i-th element of the list.
func (l *List[T]) PushAfter(i int, v T) {
	l.m.Lock()
	defer l.m.Unlock()

	l.insert(v, l.get(i))
}

// PushBefore inserts a new value before the i-th element of the list.
func (l *List[T]) PushBefore(i int, v T) {
	l.m.Lock()
	defer l.m.Unlock()

	l.insert(v, l.get(i).Prev())
}

// PopFront returns and removes the first element of the list if the list is not empty.
func (l *List[T]) PopFront() T {
	l.m.Lock()
	defer l.m.Unlock()

	v := l.root.Next().Val()
	l.remove(l.root.Next())

	return v
}

// PopBack returns and removes the last element of the list if the list is not empty.
func (l *List[T]) PopBack() T {
	l.m.Lock()
	defer l.m.Unlock()

	v := l.root.Prev().Val()
	l.remove(l.root.Prev())

	return v
}

// Remove removes the i-th element of the list if the i is less than len.
func (l *List[T]) Remove(i int) {
	l.m.Lock()
	defer l.m.Unlock()

	l.remove(l.get(i))
}

// lazyInit lazily initializes a zero List value.
func (l *List[T]) lazyInit() {
	if l.pool == nil {
		l.pool = node.NewPool[T]()
	}

	if l.root.Next() == nil {
		l.root.SetNext(&l.root)
		l.root.SetPrev(&l.root)
	}
}

// insert inserts node with value v after at, increments len.
func (l *List[T]) insert(v T, at *node.Node[T]) {
	if at == nil {
		return
	}

	n := l.pool.Pop()
	n.SetVal(v)
	at.Insert(n)
	l.len++
}

// remove removes n from list, decrements len.
func (l *List[T]) remove(n *node.Node[T]) {
	if n == nil || l.len <= 0 {
		return
	}

	n.Remove()
	l.pool.Push(n)
	l.len--
}

// get returns the i-th node or nil if the index is less than zero or greater than the len of the list.
func (l *List[T]) get(i int) *node.Node[T] {
	if l.len <= i || i < 0 {
		return nil
	}

	var (
		n    *node.Node[T]
		next func() *node.Node[T]
	)
	if i > l.len/2 {
		i = l.len - i
		n = l.root.Prev()
		next = func() *node.Node[T] {
			return n.Prev()
		}

	} else {
		n = l.root.Next()
		next = func() *node.Node[T] {
			return n.Next()
		}
	}

	for i > 0 {
		n = next()
		i--
	}

	return n
}
