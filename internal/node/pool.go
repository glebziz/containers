package node

const (
	defaultSize = 1 << 4
)

// Pool is a pool of nodes.
type Pool[T any] struct {
	free *Node[T]
	pool []Node[T]
}

// NewPool returns an initialised pool with the default capacity (16).
func NewPool[T any]() *Pool[T] {
	return NewPoolPresized[T](defaultSize)
}

// NewPoolPresized returns an initialised pool with a capacity of equal size.
func NewPoolPresized[T any](size int) *Pool[T] {
	p := Pool[T]{}
	p.init(size)

	return &p
}

// Cap returns the capacity of the pool or zero if the pool is nil.
func (p *Pool[T]) Cap() int {
	if p == nil {
		return 0
	}

	return cap(p.pool)
}

// Pop returns the first free node or the new node created node.
// If the pool is nil, nil is returned.
func (p *Pool[T]) Pop() *Node[T] {
	if p == nil {
		return nil
	}

	if p.free != nil {
		n := p.free

		p.free = p.free.next
		p.free.SetPrev(nil)
		return n
	}

	p.init(defaultSize)
	ind := len(p.pool)
	p.pool = append(p.pool, Node[T]{})
	return &p.pool[ind]
}

// Push inserts a node into the list of free nodes if pool is not nil.
func (p *Pool[T]) Push(n *Node[T]) {
	if p == nil {
		return
	}

	n.SetNext(p.free)
	n.SetPrev(nil)
	p.free = n
}

// init allocates memory for a pool of nodes.
func (p *Pool[T]) init(size int) {
	if p.pool == nil {
		p.pool = make([]Node[T], 0, size)
	}
}
