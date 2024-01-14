package omap

import (
	"testing"
)

func BenchmarkOMap_Store(b *testing.B) {
	b.Run("ordered map", func(b *testing.B) {
		b.ReportAllocs()

		m := New[int, int]()

		for i := 0; i < b.N; i++ {
			m.Store(i, i)
		}
	})

	b.Run("presized ordered map", func(b *testing.B) {
		b.ReportAllocs()

		l := NewPresized[int, int](b.N)

		for i := 0; i < b.N; i++ {
			l.Store(i, i)
		}
	})
}

func BenchmarkOMap_Delete(b *testing.B) {
	b.Run("ordered map", func(b *testing.B) {
		b.ReportAllocs()

		l := New[int, int]()

		for i := 0; i < b.N; i++ {
			l.Store(i, i)
		}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			l.Delete(i)
		}
	})

	b.Run("presized ordered map", func(b *testing.B) {
		b.ReportAllocs()

		l := NewPresized[int, int](b.N)

		for i := 0; i < b.N; i++ {
			l.Store(i, i)
		}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			l.Delete(i)
		}
	})
}

func BenchmarkOMap_Load(b *testing.B) {
	b.Run("ordered map", func(b *testing.B) {
		b.ReportAllocs()

		l := New[int, int]()

		for i := 0; i < b.N; i++ {
			l.Store(i, i)
		}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			l.Load(i)
		}
	})

	b.Run("presized ordered map", func(b *testing.B) {
		b.ReportAllocs()

		l := NewPresized[int, int](b.N)

		for i := 0; i < b.N; i++ {
			l.Store(i, i)
		}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			l.Load(i)
		}
	})
}

func BenchmarkOMap_Iter(b *testing.B) {
	b.Run("ordered map", func(b *testing.B) {
		b.ReportAllocs()

		l := New[int, int]()

		for i := 0; i < b.N; i++ {
			l.Store(i, i)
		}

		sum := 0
		b.ResetTimer()

		for it := l.Iter(); it.Next(); {
			sum += it.Val()
		}
	})

	b.Run("presized ordered map", func(b *testing.B) {
		b.ReportAllocs()

		l := NewPresized[int, int](b.N)

		for i := 0; i < b.N; i++ {
			l.Store(i, i)
		}

		sum := 0
		b.ResetTimer()

		for it := l.Iter(); it.Next(); {
			sum += it.Val()
		}
	})
}
