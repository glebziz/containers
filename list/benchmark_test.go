package list

import (
	"container/list"
	"testing"
)

func BenchmarkList_PushBack(b *testing.B) {
	b.Run("std list", func(b *testing.B) {
		b.ReportAllocs()

		l := list.New()

		for i := 0; i < b.N; i++ {
			l.PushBack(i)
		}
	})

	b.Run("list with pool", func(b *testing.B) {
		b.ReportAllocs()

		l := New[int]()

		for i := 0; i < b.N; i++ {
			l.PushBack(i)
		}
	})

	b.Run("list with presized pool", func(b *testing.B) {
		b.ReportAllocs()

		l := NewPresized[int](b.N)

		for i := 0; i < b.N; i++ {
			l.PushBack(i)
		}
	})
}

func BenchmarkList_PopBack(b *testing.B) {
	b.Run("std list", func(b *testing.B) {
		b.ReportAllocs()

		l := list.New()

		for i := 0; i < b.N; i++ {
			l.PushBack(i)
		}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			l.Remove(l.Back())
		}
	})

	b.Run("list with pool", func(b *testing.B) {
		b.ReportAllocs()

		l := New[int]()

		for i := 0; i < b.N; i++ {
			l.PushBack(i)
		}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			l.PopBack()
		}
	})

	b.Run("list with presized pool", func(b *testing.B) {
		b.ReportAllocs()

		l := NewPresized[int](b.N)

		for i := 0; i < b.N; i++ {
			l.PushBack(i)
		}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			l.PopBack()
		}
	})
}

func BenchmarkList_Iter(b *testing.B) {
	b.Run("std list", func(b *testing.B) {
		b.ReportAllocs()

		l := list.New()

		for i := 0; i < b.N; i++ {
			l.PushBack(i)
		}

		sum := 0
		b.ResetTimer()

		for e := l.Front(); e != nil; e = e.Next() {
			sum += e.Value.(int)
		}
	})

	b.Run("list with pool", func(b *testing.B) {
		b.ReportAllocs()

		l := New[int]()

		for i := 0; i < b.N; i++ {
			l.PushBack(i)
		}

		sum := 0
		b.ResetTimer()

		for it := l.Iter(); it.Next(); {
			sum += it.Val()
		}
	})

	b.Run("list with presized pool", func(b *testing.B) {
		b.ReportAllocs()

		l := NewPresized[int](b.N)

		for i := 0; i < b.N; i++ {
			l.PushBack(i)
		}

		sum := 0
		b.ResetTimer()

		for it := l.Iter(); it.Next(); {
			sum += it.Val()
		}
	})
}
