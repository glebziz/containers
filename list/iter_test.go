package list

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIter(t *testing.T) {
	const (
		N = 10
	)

	l := NewPresized[int](N)
	for i := 0; i < N; i++ {
		l.PushBack(i)
	}

	i := 0
	it := l.Iter()

	for it.Next() {
		require.Equal(t, i, it.Val())
		i++
	}

	require.Equal(t, N, i)
}

func TestRIter(t *testing.T) {
	const (
		N = 10
	)

	l := NewPresized[int](N)
	for i := 0; i < N; i++ {
		l.PushBack(i)
	}

	i := N
	it := l.RIter()

	for it.Next() {
		require.Equal(t, i-1, it.Val())
		i--
	}

	require.Zero(t, i)
}
