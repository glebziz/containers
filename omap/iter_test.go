package omap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIter(t *testing.T) {
	const (
		N = 10
	)

	l := NewPresized[int, int](N)
	for i := 0; i < N; i++ {
		l.Store(i, i)
	}

	i := 0
	it := l.Iter()

	for it.Next() {
		require.Equal(t, i, it.Val())
		i++
	}

	require.Equal(t, N, i)
}
