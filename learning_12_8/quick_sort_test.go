package main

import (
	"benchForSort/additional"
	"benchForSort/mySorts"
	"testing"
)

// 12.7 алгоритм быстрой сортировки
func BenchmarkQuickSort(b *testing.B) {
	b.Run("small arrays quick", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := additional.GenerateSlice(10, 10)
			b.StartTimer()
			mySorts.QuickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays quick", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := additional.GenerateSlice(1000, 1000)
			b.StartTimer()
			mySorts.QuickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays quick", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := additional.GenerateSlice(100000, 100000)
			b.StartTimer()
			mySorts.QuickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("small arrays quick 12.8.4", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := additional.GenerateSlice(100, 3)
			b.StartTimer()
			mySorts.QuickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays quick 12.8.4", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := additional.GenerateSlice(100, 300000)
			b.StartTimer()
			mySorts.QuickSort(ar)
			b.StopTimer()
		}
	})
}
