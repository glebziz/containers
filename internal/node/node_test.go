package node

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNode_SetVal(t *testing.T) {
	for _, tc := range []struct {
		name    string
		node    *Node[int]
		val     int
		expNode *Node[int]
	}{
		{
			name: "nil node",
			val:  10,
		},
		{
			name: "not nil node",
			node: &Node[int]{},
			val:  10,
			expNode: &Node[int]{
				val: 10,
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.node.SetVal(tc.val)
			require.Equal(t, tc.expNode, tc.node)
		})
	}
}

func TestNode_SetNext(t *testing.T) {
	for _, tc := range []struct {
		name    string
		node    *Node[int]
		next    *Node[int]
		expNode *Node[int]
	}{
		{
			name: "nil node and nil next node",
		},
		{
			name: "nil node and not nil next node",
			next: &Node[int]{
				val: 10,
			},
		},
		{
			name: "not nil node and nil next node",
			node: &Node[int]{
				next: &Node[int]{},
				prev: &Node[int]{},
			},
			expNode: &Node[int]{
				prev: &Node[int]{},
			},
		},
		{
			name: "not nil node and not nil next node",
			node: &Node[int]{
				next: &Node[int]{},
				prev: &Node[int]{},
			},
			next: &Node[int]{
				val: 10,
			},
			expNode: &Node[int]{
				next: &Node[int]{
					val: 10,
				},
				prev: &Node[int]{},
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.node.SetNext(tc.next)
			require.Equal(t, tc.expNode, tc.node)
		})
	}
}

func TestNode_SetPrev(t *testing.T) {
	for _, tc := range []struct {
		name    string
		node    *Node[int]
		prev    *Node[int]
		expNode *Node[int]
	}{
		{
			name: "nil node and nil prev node",
		},
		{
			name: "nil node and not nil prev node",
			prev: &Node[int]{
				val: 10,
			},
		},
		{
			name: "not nil node and nil prev node",
			node: &Node[int]{
				next: &Node[int]{},
				prev: &Node[int]{},
			},
			expNode: &Node[int]{
				next: &Node[int]{},
			},
		},
		{
			name: "not nil node and not nil next node",
			node: &Node[int]{
				next: &Node[int]{},
				prev: &Node[int]{},
			},
			prev: &Node[int]{
				val: 10,
			},
			expNode: &Node[int]{
				next: &Node[int]{},
				prev: &Node[int]{
					val: 10,
				},
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.node.SetPrev(tc.prev)
			require.Equal(t, tc.expNode, tc.node)
		})
	}
}

func TestNode_Val(t *testing.T) {
	for _, tc := range []struct {
		name   string
		node   *Node[int]
		expVal int
	}{
		{
			name: "nil node",
		},
		{
			name: "not nil node",
			node: &Node[int]{
				val: 10,
			},
			expVal: 10,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			val := tc.node.Val()
			require.Equal(t, tc.expVal, val)
		})
	}
}

func TestNode_Next(t *testing.T) {
	for _, tc := range []struct {
		name    string
		node    *Node[int]
		expNext *Node[int]
	}{
		{
			name: "nil node",
		},
		{
			name: "not nil node and nil next node",
			node: &Node[int]{
				prev: &Node[int]{},
			},
		},
		{
			name: "not nil node and not nil next node",
			node: &Node[int]{
				next: &Node[int]{
					val: 10,
				},
				prev: &Node[int]{},
			},
			expNext: &Node[int]{
				val: 10,
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			next := tc.node.Next()
			require.Equal(t, tc.expNext, next)
		})
	}
}

func TestNode_Prev(t *testing.T) {
	for _, tc := range []struct {
		name    string
		node    *Node[int]
		expPrev *Node[int]
	}{
		{
			name: "nil node",
		},
		{
			name: "not nil node and nil prev node",
			node: &Node[int]{
				next: &Node[int]{},
			},
		},
		{
			name: "not nil node and not nil next node",
			node: &Node[int]{
				next: &Node[int]{},
				prev: &Node[int]{
					val: 10,
				},
			},
			expPrev: &Node[int]{
				val: 10,
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			prev := tc.node.Prev()
			require.Equal(t, tc.expPrev, prev)
		})
	}
}

func TestNode_Insert(t *testing.T) {
	for _, tc := range []struct {
		name      string
		node      func() *Node[int]
		newNode   *Node[int]
		checkNode func(t *testing.T, n *Node[int])
	}{
		{
			name: "nil node",
			node: func() *Node[int] {
				return nil
			},
			newNode: &Node[int]{},
			checkNode: func(t *testing.T, n *Node[int]) {
				require.Nil(t, n)
			},
		},
		{
			name: "between two nodes",
			node: func() *Node[int] {
				f := &Node[int]{
					val: 1,
				}
				s := &Node[int]{
					val:  2,
					prev: f,
				}

				f.next = s

				return f
			},
			newNode: &Node[int]{
				val: 3,
			},
			checkNode: func(t *testing.T, n *Node[int]) {
				require.Equal(t, 1, n.val)
				require.Equal(t, 3, n.next.val)
				require.Equal(t, 2, n.next.next.val)
				require.Equal(t, n, n.next.prev)
				require.Equal(t, n.next, n.next.next.prev)
			},
		},
		{
			name: "after one node",
			node: func() *Node[int] {
				f := &Node[int]{
					val: 1,
				}

				return f
			},
			newNode: &Node[int]{
				val: 2,
			},
			checkNode: func(t *testing.T, n *Node[int]) {
				require.Equal(t, 1, n.val)
				require.Equal(t, 2, n.next.val)
				require.Equal(t, n, n.next.prev)
				require.Nil(t, n.next.next)
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			n := tc.node()
			n.Insert(tc.newNode)
			tc.checkNode(t, n)
		})
	}
}

func TestNode_Remove(t *testing.T) {
	for _, tc := range []struct {
		name      string
		node      func() (first *Node[int], deleted *Node[int])
		checkNode func(t *testing.T, n *Node[int])
	}{
		{
			name: "nil node",
			node: func() (first *Node[int], deleted *Node[int]) {
				return nil, nil
			},
			checkNode: func(t *testing.T, n *Node[int]) {
				require.Nil(t, n)
			},
		},
		{
			name: "between two nodes",
			node: func() (first *Node[int], deleted *Node[int]) {
				first = &Node[int]{
					val: 1,
				}
				second := &Node[int]{
					val:  2,
					prev: first,
				}
				third := &Node[int]{
					val:  3,
					prev: second,
				}

				first.next = second
				second.next = third

				return first, second
			},
			checkNode: func(t *testing.T, n *Node[int]) {
				require.Equal(t, 1, n.val)
				require.Equal(t, 3, n.next.val)
				require.Equal(t, n, n.next.prev)
				require.Nil(t, n.next.next)
			},
		},
		{
			name: "first node",
			node: func() (first *Node[int], deleted *Node[int]) {
				first = &Node[int]{
					val: 1,
				}
				second := &Node[int]{
					val:  2,
					prev: first,
				}

				first.next = second

				return second, first
			},
			checkNode: func(t *testing.T, n *Node[int]) {
				require.Equal(t, 2, n.val)
				require.Nil(t, n.next)
				require.Nil(t, n.prev)
			},
		},
		{
			name: "last node",
			node: func() (first *Node[int], deleted *Node[int]) {
				first = &Node[int]{
					val: 1,
				}
				second := &Node[int]{
					val:  2,
					prev: first,
				}

				first.next = second

				return first, second
			},
			checkNode: func(t *testing.T, n *Node[int]) {
				require.Equal(t, 1, n.val)
				require.Nil(t, n.next)
				require.Nil(t, n.prev)
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			f, n := tc.node()
			n.Remove()
			tc.checkNode(t, f)
		})
	}
}
