package iter

import "github.com/glebziz/containers/internal/node"

// Direction is the direction of the iterator.
type Direction bool

const (
	// ForwardDir is the forward direction.
	ForwardDir = Direction(true)
	// ReverseDir is the reverse direction.
	ReverseDir = Direction(false)
)

// Iter is an iterator that supports iterating over the values of container types.
type Iter[T any] struct {
	dir  Direction
	c    *node.Node[T]
	stop *node.Node[T]
}

// New returns an initialised iterator.
func New[T any](n *node.Node[T], dir Direction) *Iter[T] {
	return &Iter[T]{
		dir:  dir,
		c:    n,
		stop: n,
	}
}

// Next selects the next node and returns true if it exists.
// Otherwise, it returns false.
func (i *Iter[T]) Next() bool {
	var next *node.Node[T]
	switch i.dir {
	case ForwardDir:
		next = i.c.Next()
	case ReverseDir:
		next = i.c.Prev()
	}

	if next == nil || next == i.stop {
		return false
	}

	i.c = next
	return true
}

// Val returns the value of the current node.
func (i *Iter[T]) Val() T {
	return i.c.Val()
}
