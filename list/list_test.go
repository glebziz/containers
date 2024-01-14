package list

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/glebziz/containers/internal/node"
)

func TestList_Front(t *testing.T) {
	for _, tc := range []struct {
		name   string
		l      func() *List[int]
		expVal int
	}{
		{
			name: "empty list",
			l: func() *List[int] {
				return New[int]()
			},
			expVal: 0,
		},
		{
			name: "not empty list",
			l: func() *List[int] {
				l := New[int]()
				next := node.Node[int]{}

				next.SetVal(10)
				l.root.SetNext(&next)

				return l
			},
			expVal: 10,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l := tc.l()
			v := l.Front()
			require.Equal(t, tc.expVal, v)

			v = l.Front()
			require.Equal(t, tc.expVal, v)
		})
	}
}

func TestList_Back(t *testing.T) {
	for _, tc := range []struct {
		name   string
		l      func() *List[int]
		expVal int
	}{
		{
			name: "empty list",
			l: func() *List[int] {
				return New[int]()
			},
			expVal: 0,
		},
		{
			name: "not empty list",
			l: func() *List[int] {
				l := New[int]()
				next := node.Node[int]{}

				next.SetVal(10)
				l.root.SetPrev(&next)

				return l
			},
			expVal: 10,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l := tc.l()
			v := l.Back()
			require.Equal(t, tc.expVal, v)

			v = l.Back()
			require.Equal(t, tc.expVal, v)
		})
	}
}

func TestList_Get(t *testing.T) {
	for _, tc := range []struct {
		name   string
		i      int
		l      func() *List[int]
		expVal int
	}{
		{
			name: "empty list",
			i:    0,
			l: func() *List[int] {
				return New[int]()
			},
			expVal: 0,
		},
		{
			name: "not empty list",
			i:    0,
			l: func() *List[int] {
				l := New[int]()
				next := node.Node[int]{}

				next.SetVal(10)
				l.root.SetNext(&next)
				l.len = 1

				return l
			},
			expVal: 10,
		},
		{
			name: "list with more values, i less than median",
			i:    1,
			l: func() *List[int] {
				l := New[int]()
				first := node.Node[int]{}
				second := node.Node[int]{}
				third := node.Node[int]{}
				fourth := node.Node[int]{}

				first.SetVal(10)
				second.SetVal(20)
				third.SetVal(30)
				fourth.SetVal(40)

				l.root.SetNext(&first)
				l.root.SetPrev(&fourth)
				first.SetNext(&second)
				second.SetNext(&third)
				third.SetNext(&fourth)
				second.SetPrev(&first)
				third.SetPrev(&second)
				fourth.SetPrev(&third)

				l.len = 4

				return l
			},
			expVal: 20,
		},
		{
			name: "list with more values, i greater than median",
			i:    3,
			l: func() *List[int] {
				l := New[int]()
				first := node.Node[int]{}
				second := node.Node[int]{}
				third := node.Node[int]{}
				fourth := node.Node[int]{}

				first.SetVal(10)
				second.SetVal(20)
				third.SetVal(30)
				fourth.SetVal(40)

				l.root.SetNext(&first)
				l.root.SetPrev(&fourth)
				first.SetNext(&second)
				second.SetNext(&third)
				third.SetNext(&fourth)
				second.SetPrev(&first)
				third.SetPrev(&second)
				fourth.SetPrev(&third)

				l.len = 4

				return l
			},
			expVal: 30,
		},
		{
			name: "index out of range",
			i:    3,
			l: func() *List[int] {
				l := New[int]()
				first := node.Node[int]{}
				second := node.Node[int]{}

				first.SetVal(10)
				second.SetVal(20)

				l.root.SetNext(&first)
				l.root.SetPrev(&second)
				first.SetNext(&second)
				second.SetPrev(&first)

				l.len = 2

				return l
			},
			expVal: 0,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l := tc.l()
			v := l.Get(tc.i)
			require.Equal(t, tc.expVal, v)

			v = l.Get(tc.i)
			require.Equal(t, tc.expVal, v)
		})
	}
}

func TestList_PushFront(t *testing.T) {
	for _, tc := range []struct {
		name      string
		v         int
		l         func() *List[int]
		checkList func(t *testing.T, l *List[int])
	}{
		{
			name: "empty list",
			v:    10,
			l: func() *List[int] {
				return &List[int]{}
			},
			checkList: func(t *testing.T, l *List[int]) {
				require.Equal(t, 1, l.Len())
				require.Equal(t, 10, l.root.Next().Val())
				require.Equal(t, l.root.Next(), l.root.Prev())
			},
		},
		{
			name: "not empty list",
			v:    10,
			l: func() *List[int] {
				l := NewPresized[int](1)
				next := node.Node[int]{}

				next.SetVal(1)
				next.SetNext(&l.root)
				next.SetPrev(&l.root)
				l.root.SetNext(&next)
				l.root.SetPrev(&next)
				l.len = 1

				return l
			},
			checkList: func(t *testing.T, l *List[int]) {
				require.Equal(t, 2, l.Len())
				require.Equal(t, 10, l.root.Next().Val())
				require.Equal(t, 1, l.root.Prev().Val())
				require.Equal(t, l.root.Next().Next(), l.root.Prev())
				require.Equal(t, l.root.Prev().Prev(), l.root.Next())
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l := tc.l()
			l.PushFront(tc.v)
			tc.checkList(t, l)
		})
	}
}

func TestList_PushBack(t *testing.T) {
	for _, tc := range []struct {
		name      string
		v         int
		l         func() *List[int]
		checkList func(t *testing.T, l *List[int])
	}{
		{
			name: "empty list",
			v:    10,
			l: func() *List[int] {
				return &List[int]{}
			},
			checkList: func(t *testing.T, l *List[int]) {
				require.Equal(t, 1, l.Len())
				require.Equal(t, 10, l.root.Next().Val())
				require.Equal(t, l.root.Next(), l.root.Prev())
			},
		},
		{
			name: "not empty list",
			v:    10,
			l: func() *List[int] {
				l := NewPresized[int](1)
				next := node.Node[int]{}

				next.SetVal(1)
				next.SetNext(&l.root)
				next.SetPrev(&l.root)
				l.root.SetNext(&next)
				l.root.SetPrev(&next)
				l.len = 1

				return l
			},
			checkList: func(t *testing.T, l *List[int]) {
				require.Equal(t, 2, l.Len())
				require.Equal(t, 1, l.root.Next().Val())
				require.Equal(t, 10, l.root.Prev().Val())
				require.Equal(t, l.root.Next().Next(), l.root.Prev())
				require.Equal(t, l.root.Prev().Prev(), l.root.Next())
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l := tc.l()
			l.PushBack(tc.v)
			tc.checkList(t, l)
		})
	}
}

func TestList_PushAfter(t *testing.T) {
	for _, tc := range []struct {
		name      string
		i         int
		v         int
		l         func() *List[int]
		checkList func(t *testing.T, l *List[int])
	}{
		{
			name: "empty list",
			i:    0,
			v:    10,
			l: func() *List[int] {
				return &List[int]{}
			},
			checkList: func(t *testing.T, l *List[int]) {
				require.Zero(t, l.Len())
				require.Equal(t, l.root.Next(), l.root.Prev())
				require.Nil(t, l.root.Next())
			},
		},
		{
			name: "not empty list",
			i:    0,
			v:    10,
			l: func() *List[int] {
				l := NewPresized[int](1)
				next := node.Node[int]{}

				next.SetVal(1)
				next.SetNext(&l.root)
				next.SetPrev(&l.root)
				l.root.SetNext(&next)
				l.root.SetPrev(&next)
				l.len = 1

				return l
			},
			checkList: func(t *testing.T, l *List[int]) {
				require.Equal(t, 2, l.Len())
				require.Equal(t, 1, l.root.Next().Val())
				require.Equal(t, 10, l.root.Prev().Val())
				require.Equal(t, l.root.Next().Next(), l.root.Prev())
				require.Equal(t, l.root.Prev().Prev(), l.root.Next())
			},
		},
		{
			name: "index out of range",
			i:    10,
			v:    10,
			l: func() *List[int] {
				l := New[int]()
				next := node.Node[int]{}

				next.SetVal(1)
				next.SetNext(&l.root)
				next.SetPrev(&l.root)
				l.root.SetNext(&next)
				l.root.SetPrev(&next)
				l.len = 1

				return l
			},
			checkList: func(t *testing.T, l *List[int]) {
				require.Equal(t, 1, l.Len())
				require.Equal(t, 1, l.root.Next().Val())
				require.Equal(t, l.root.Next(), l.root.Prev())
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l := tc.l()
			l.PushAfter(tc.i, tc.v)
			tc.checkList(t, l)
		})
	}
}

func TestList_PushBefore(t *testing.T) {
	for _, tc := range []struct {
		name      string
		i         int
		v         int
		l         func() *List[int]
		checkList func(t *testing.T, l *List[int])
	}{
		{
			name: "empty list",
			i:    0,
			v:    10,
			l: func() *List[int] {
				return &List[int]{}
			},
			checkList: func(t *testing.T, l *List[int]) {
				require.Zero(t, l.Len())
				require.Equal(t, l.root.Next(), l.root.Prev())
				require.Nil(t, l.root.Next())
			},
		},
		{
			name: "not empty list",
			i:    0,
			v:    10,
			l: func() *List[int] {
				l := NewPresized[int](1)
				next := node.Node[int]{}

				next.SetVal(1)
				next.SetNext(&l.root)
				next.SetPrev(&l.root)
				l.root.SetNext(&next)
				l.root.SetPrev(&next)
				l.len = 1

				return l
			},
			checkList: func(t *testing.T, l *List[int]) {
				require.Equal(t, 2, l.Len())
				require.Equal(t, 10, l.root.Next().Val())
				require.Equal(t, 1, l.root.Prev().Val())
				require.Equal(t, l.root.Next().Next(), l.root.Prev())
				require.Equal(t, l.root.Prev().Prev(), l.root.Next())
			},
		},
		{
			name: "index out of range",
			i:    10,
			v:    10,
			l: func() *List[int] {
				l := New[int]()
				next := node.Node[int]{}

				next.SetVal(1)
				next.SetNext(&l.root)
				next.SetPrev(&l.root)
				l.root.SetNext(&next)
				l.root.SetPrev(&next)
				l.len = 1

				return l
			},
			checkList: func(t *testing.T, l *List[int]) {
				require.Equal(t, 1, l.Len())
				require.Equal(t, 1, l.root.Next().Val())
				require.Equal(t, l.root.Next(), l.root.Prev())
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l := tc.l()
			l.PushBefore(tc.i, tc.v)
			tc.checkList(t, l)
		})
	}
}

func TestList_PopFront(t *testing.T) {
	for _, tc := range []struct {
		name      string
		l         func() *List[int]
		expVal    int
		checkList func(t *testing.T, l *List[int])
	}{
		{
			name: "empty list",
			l: func() *List[int] {
				l := New[int]()
				l.lazyInit()
				return l
			},
			expVal: 0,
			checkList: func(t *testing.T, l *List[int]) {
				require.Zero(t, l.Len())
				require.Equal(t, l.root.Next(), l.root.Prev())
				require.Equal(t, l.root.Next(), &l.root)
			},
		},
		{
			name: "not empty list",
			l: func() *List[int] {
				l := New[int]()
				next := node.Node[int]{}

				next.SetVal(10)
				next.SetNext(&l.root)
				next.SetPrev(&l.root)
				l.root.SetNext(&next)
				l.root.SetPrev(&next)
				l.len = 1

				return l
			},
			expVal: 10,
			checkList: func(t *testing.T, l *List[int]) {
				require.Zero(t, l.Len())
				require.Equal(t, l.root.Next(), l.root.Prev())
				require.Equal(t, l.root.Next(), &l.root)
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l := tc.l()
			v := l.PopFront()
			require.Equal(t, tc.expVal, v)
			tc.checkList(t, l)
		})
	}
}

func TestList_PopBack(t *testing.T) {
	for _, tc := range []struct {
		name      string
		l         func() *List[int]
		expVal    int
		checkList func(t *testing.T, l *List[int])
	}{
		{
			name: "empty list",
			l: func() *List[int] {
				return New[int]()
			},
			expVal: 0,
			checkList: func(t *testing.T, l *List[int]) {
				require.Zero(t, l.Len())
				require.Equal(t, l.root.Next(), l.root.Prev())
				require.Nil(t, l.root.Next())
			},
		},
		{
			name: "not empty list",
			l: func() *List[int] {
				l := New[int]()
				next := node.Node[int]{}

				next.SetVal(10)
				next.SetNext(&l.root)
				next.SetPrev(&l.root)
				l.root.SetNext(&next)
				l.root.SetPrev(&next)
				l.len = 1

				return l
			},
			expVal: 10,
			checkList: func(t *testing.T, l *List[int]) {
				require.Zero(t, l.Len())
				require.Equal(t, l.root.Next(), l.root.Prev())
				require.Equal(t, l.root.Next(), &l.root)
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l := tc.l()
			v := l.PopBack()
			require.Equal(t, tc.expVal, v)
			tc.checkList(t, l)
		})
	}
}

func TestList_Remove(t *testing.T) {
	for _, tc := range []struct {
		name      string
		i         int
		l         func() *List[int]
		checkList func(t *testing.T, l *List[int])
	}{
		{
			name: "empty list",
			i:    0,
			l: func() *List[int] {
				l := New[int]()
				l.lazyInit()
				return l
			},
			checkList: func(t *testing.T, l *List[int]) {
				require.Zero(t, l.Len())
				require.Equal(t, l.root.Next(), l.root.Prev())
				require.Equal(t, l.root.Next(), &l.root)
			},
		},
		{
			name: "not empty list",
			i:    0,
			l: func() *List[int] {
				l := New[int]()
				next := node.Node[int]{}

				next.SetVal(10)
				next.SetNext(&l.root)
				next.SetPrev(&l.root)
				l.root.SetNext(&next)
				l.root.SetPrev(&next)
				l.len = 1

				return l
			},
			checkList: func(t *testing.T, l *List[int]) {
				require.Zero(t, l.Len())
				require.Equal(t, l.root.Next(), l.root.Prev())
				require.Equal(t, l.root.Next(), &l.root)
			},
		},
		{
			name: "index out of range",
			i:    10,
			l: func() *List[int] {
				l := New[int]()
				next := node.Node[int]{}

				next.SetVal(10)
				next.SetNext(&l.root)
				next.SetPrev(&l.root)
				l.root.SetNext(&next)
				l.root.SetPrev(&next)
				l.len = 1

				return l
			},
			checkList: func(t *testing.T, l *List[int]) {
				require.Equal(t, 1, l.Len())
				require.Equal(t, l.root.Next(), l.root.Prev())
				require.Equal(t, l.root.Next().Next(), &l.root)
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l := tc.l()
			l.Remove(tc.i)
			tc.checkList(t, l)
		})
	}
}
