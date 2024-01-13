package node

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPool(t *testing.T) {
	p := NewPool[int]()
	require.Equal(t, defaultSize, cap(p.pool))
}

func TestNewPoolPresized(t *testing.T) {
	const (
		size = 100
	)

	p := NewPoolPresized[int](size)
	require.Equal(t, size, cap(p.pool))
}

func TestPool_Cap(t *testing.T) {
	for _, tc := range []struct {
		name string
		p    *Pool[int]
		cap  int
	}{
		{
			name: "nil pool",
		},
		{
			name: "empty pool",
			p:    &Pool[int]{},
		},
		{
			name: "presized pool",
			p:    NewPool[int](),
			cap:  defaultSize,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			c := tc.p.Cap()
			require.Equal(t, tc.cap, c)
		})
	}
}

func TestPool_Pop(t *testing.T) {
	for _, tc := range []struct {
		name    string
		p       *Pool[int]
		expP    *Pool[int]
		expNode *Node[int]
	}{
		{
			name: "nil pool",
		},
		{
			name: "empty pool",
			p:    NewPool[int](),
			expP: &Pool[int]{
				pool: []Node[int]{{}},
			},
			expNode: &Node[int]{},
		},
		{
			name: "pool with free nodes",
			p: &Pool[int]{
				free: &Node[int]{
					val: 10,
				},
			},
			expP: &Pool[int]{},
			expNode: &Node[int]{
				val: 10,
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			n := tc.p.Pop()

			require.Equal(t, tc.expP, tc.p)
			require.Equal(t, tc.expNode, n)
		})
	}
}

func TestPool_Push(t *testing.T) {
	t.Parallel()

	p := NewPool[int]()

	n := &Node[int]{}

	p.Push(n)

	require.Equal(t, n, p.free)
}
