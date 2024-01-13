package omap

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/glebziz/containers/internal/node"
)

func TestOMap_Store(t *testing.T) {
	for _, tc := range []struct {
		name     string
		k        int
		v        int
		m        func() *OMap[int, int]
		checkMap func(t *testing.T, m *OMap[int, int])
	}{
		{
			name: "empty map",
			k:    20,
			v:    10,
			m: func() *OMap[int, int] {
				return &OMap[int, int]{}
			},
			checkMap: func(t *testing.T, m *OMap[int, int]) {
				require.Equal(t, 1, m.Len())
				require.Equal(t, 10, m.root.Next().Val())
				require.Equal(t, m.root.Next(), m.data[20])
				require.Equal(t, m.root.Next(), m.root.Prev())
			},
		},
		{
			name: "not empty map",
			k:    20,
			v:    10,
			m: func() *OMap[int, int] {
				m := NewPresized[int, int](1)
				next := node.Node[int]{}

				next.SetVal(1)
				next.SetNext(&m.root)
				next.SetPrev(&m.root)
				m.root.SetNext(&next)
				m.root.SetPrev(&next)
				m.data[2] = &next

				return m
			},
			checkMap: func(t *testing.T, m *OMap[int, int]) {
				require.Equal(t, 2, m.Len())
				require.Equal(t, 1, m.root.Next().Val())
				require.Equal(t, 10, m.root.Prev().Val())
				require.Equal(t, m.root.Prev(), m.data[20])
				require.Equal(t, m.root.Next().Next(), m.root.Prev())
				require.Equal(t, m.root.Prev().Prev(), m.root.Next())
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			m := tc.m()
			m.Store(tc.k, tc.v)
			tc.checkMap(t, m)
		})
	}
}

func TestOMap_Load(t *testing.T) {
	for _, tc := range []struct {
		name   string
		k      int
		m      func() *OMap[int, int]
		expVal int
		expOk  bool
	}{
		{
			name: "empty map",
			k:    20,
			m: func() *OMap[int, int] {
				return &OMap[int, int]{}
			},
			expVal: 0,
			expOk:  false,
		},
		{
			name: "map with more values",
			k:    20,
			m: func() *OMap[int, int] {
				m := New[int, int]()
				first := node.Node[int]{}
				second := node.Node[int]{}
				third := node.Node[int]{}

				first.SetVal(1)
				second.SetVal(10)
				third.SetVal(100)

				m.root.SetNext(&first)
				m.root.SetPrev(&third)
				first.SetNext(&second)
				second.SetNext(&third)
				third.SetNext(&m.root)
				second.SetPrev(&first)
				third.SetPrev(&second)
				m.data[2] = &first
				m.data[20] = &second
				m.data[200] = &third

				return m
			},
			expVal: 10,
			expOk:  true,
		},
		{
			name: "key not found",
			k:    20,
			m: func() *OMap[int, int] {
				m := New[int, int]()
				first := node.Node[int]{}

				first.SetVal(1)

				m.root.SetNext(&first)
				m.root.SetPrev(&first)
				m.data[2] = &first

				return m
			},
			expVal: 0,
			expOk:  false,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			m := tc.m()
			v, ok := m.Load(tc.k)
			require.Equal(t, tc.expOk, ok)
			require.Equal(t, tc.expVal, v)

			v, ok = m.Load(tc.k)
			require.Equal(t, tc.expOk, ok)
			require.Equal(t, tc.expVal, v)
		})
	}
}

func TestOMap_Delete(t *testing.T) {
	for _, tc := range []struct {
		name     string
		k        int
		m        func() *OMap[int, int]
		checkMap func(t *testing.T, l *OMap[int, int])
	}{
		{
			name: "empty map",
			k:    20,
			m: func() *OMap[int, int] {
				return New[int, int]()
			},
			checkMap: func(t *testing.T, m *OMap[int, int]) {
				require.Zero(t, m.Len())
				require.Equal(t, m.root.Next(), m.root.Prev())
				require.Nil(t, m.root.Next())
			},
		},
		{
			name: "not empty map",
			k:    20,
			m: func() *OMap[int, int] {
				m := New[int, int]()
				first := node.Node[int]{}
				second := node.Node[int]{}
				third := node.Node[int]{}

				first.SetVal(1)
				second.SetVal(10)
				third.SetVal(100)

				m.root.SetNext(&first)
				m.root.SetPrev(&third)
				first.SetNext(&second)
				second.SetNext(&third)
				third.SetNext(&m.root)
				second.SetPrev(&first)
				third.SetPrev(&second)
				m.data[2] = &first
				m.data[20] = &second
				m.data[200] = &third

				return m
			},
			checkMap: func(t *testing.T, m *OMap[int, int]) {
				require.Equal(t, 2, m.Len())
				require.Equal(t, 100, m.root.Next().Next().Val())
			},
		},
		{
			name: "key not found",
			k:    20,
			m: func() *OMap[int, int] {
				m := New[int, int]()
				first := node.Node[int]{}

				first.SetVal(1)

				m.root.SetNext(&first)
				m.root.SetPrev(&first)
				m.data[2] = &first

				return m
			},
			checkMap: func(t *testing.T, m *OMap[int, int]) {
				require.Equal(t, 1, m.Len())
				require.Equal(t, m.root.Next(), m.root.Prev())
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l := tc.m()
			l.Delete(tc.k)
			tc.checkMap(t, l)
		})
	}
}
