package node

// Node is an element of container structures.
type Node[T any] struct {
	val  T
	next *Node[T]
	prev *Node[T]
}

// SetVal sets the value if the node is not nil.
func (n *Node[T]) SetVal(v T) {
	if n == nil {
		return
	}

	n.val = v
}

// SetNext sets the next node pointer if the node is not nil.
func (n *Node[T]) SetNext(next *Node[T]) {
	if n == nil {
		return
	}

	n.next = next
}

// SetPrev sets the previous node pointer if the node is not nil.
func (n *Node[T]) SetPrev(prev *Node[T]) {
	if n == nil {
		return
	}

	n.prev = prev
}

// Val returns the value of the node or zero value if the node is nil.
func (n *Node[T]) Val() T {
	if n == nil {
		var v T
		return v
	}

	return n.val
}

// Next returns the next node or nil if the node is nil.
func (n *Node[T]) Next() *Node[T] {
	if n == nil {
		return nil
	}

	return n.next
}

// Prev returns the previous node or nil if the node is nil.
func (n *Node[T]) Prev() *Node[T] {
	if n == nil {
		return nil
	}

	return n.prev
}

// Insert inserts a new node after the current node if it is not nil.
func (n *Node[T]) Insert(new *Node[T]) {
	if n == nil {
		return
	}

	new.SetPrev(n)
	new.SetNext(n.Next())
	n.Next().SetPrev(new)
	n.SetNext(new)
}

// Remove removes the current node if is not nil.
func (n *Node[T]) Remove() {
	if n == nil {
		return
	}

	n.Next().SetPrev(n.Prev())
	n.Prev().SetNext(n.Next())
}
