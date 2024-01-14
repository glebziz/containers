package omap

import (
	"sync"
	"testing"
)

func BenchmarkOMap_Store(b *testing.B) {
	b.Run("sync map", func(b *testing.B) {
		b.ReportAllocs()

		var m sync.Map

		for i := 0; i < b.N; i++ {
			m.Store(i, i)
		}
	})

	b.Run("ordered map", func(b *testing.B) {
		b.ReportAllocs()

		m := New[int, int]()

		for i := 0; i < b.N; i++ {
			m.Store(i, i)
		}
	})

	b.Run("presized ordered map", func(b *testing.B) {
		b.ReportAllocs()

		m := NewPresized[int, int](b.N)

		for i := 0; i < b.N; i++ {
			m.Store(i, i)
		}
	})
}

func BenchmarkOMap_Delete(b *testing.B) {
	b.Run("sync map", func(b *testing.B) {
		b.ReportAllocs()

		var m sync.Map

		for i := 0; i < b.N; i++ {
			m.Store(i, i)
		}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			m.Delete(i)
		}
	})

	b.Run("ordered map", func(b *testing.B) {
		b.ReportAllocs()

		m := New[int, int]()

		for i := 0; i < b.N; i++ {
			m.Store(i, i)
		}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			m.Delete(i)
		}
	})

	b.Run("presized ordered map", func(b *testing.B) {
		b.ReportAllocs()

		m := NewPresized[int, int](b.N)

		for i := 0; i < b.N; i++ {
			m.Store(i, i)
		}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			m.Delete(i)
		}
	})
}

func BenchmarkOMap_Load(b *testing.B) {
	b.Run("sync map", func(b *testing.B) {
		b.ReportAllocs()

		var m sync.Map

		for i := 0; i < b.N; i++ {
			m.Store(i, i)
		}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			m.Load(i)
		}
	})

	b.Run("ordered map", func(b *testing.B) {
		b.ReportAllocs()

		m := New[int, int]()

		for i := 0; i < b.N; i++ {
			m.Store(i, i)
		}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			m.Load(i)
		}
	})

	b.Run("presized ordered map", func(b *testing.B) {
		b.ReportAllocs()

		m := NewPresized[int, int](b.N)

		for i := 0; i < b.N; i++ {
			m.Store(i, i)
		}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			m.Load(i)
		}
	})
}

func BenchmarkOMap_Iter(b *testing.B) {
	b.Run("sync map", func(b *testing.B) {
		b.ReportAllocs()

		var m sync.Map

		for i := 0; i < b.N; i++ {
			m.Store(i, i)
		}

		sum := 0
		b.ResetTimer()

		m.Range(func(key, value any) bool {
			sum += value.(int)
			return true
		})
	})

	b.Run("ordered map", func(b *testing.B) {
		b.ReportAllocs()

		m := New[int, int]()

		for i := 0; i < b.N; i++ {
			m.Store(i, i)
		}

		sum := 0
		b.ResetTimer()

		for it := m.Iter(); it.Next(); {
			sum += it.Val()
		}
	})

	b.Run("presized ordered map", func(b *testing.B) {
		b.ReportAllocs()

		m := NewPresized[int, int](b.N)

		for i := 0; i < b.N; i++ {
			m.Store(i, i)
		}

		sum := 0
		b.ResetTimer()

		for it := m.Iter(); it.Next(); {
			sum += it.Val()
		}
	})
}
