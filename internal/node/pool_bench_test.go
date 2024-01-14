package node

import (
	"testing"
)

func BenchmarkPool_Pop(b *testing.B) {
	for _, tc := range []struct {
		name string
		p    *Pool[int]
	}{
		{
			name: "pool",
			p:    NewPool[int](),
		},
		{
			name: "presized pool",
			p:    NewPoolPresized[int](b.N),
		},
	} {
		tc := tc
		b.Run(tc.name, func(b *testing.B) {
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				tc.p.Pop()
			}
		})
	}
}

func BenchmarkPool_Push(b *testing.B) {
	var (
		n = Node[int]{}
	)

	for _, tc := range []struct {
		name string
		p    *Pool[int]
	}{
		{
			name: "pool",
			p:    NewPool[int](),
		},
		{
			name: "presized pool",
			p:    NewPoolPresized[int](b.N),
		},
	} {
		tc := tc
		b.Run(tc.name, func(b *testing.B) {
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				tc.p.Push(&n)
			}
		})
	}
}
