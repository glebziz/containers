package omap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIter(t *testing.T) {
	const (
		N = 10
	)

	m := NewPresized[int, int](N)

	it := m.Iter()
	require.False(t, it.Next())

	for i := 0; i < N; i++ {
		m.Store(i, i)
	}

	i := 0
	it = m.Iter()

	for it.Next() {
		require.Equal(t, i, it.Val())
		i++
	}

	require.Equal(t, N, i)
}
