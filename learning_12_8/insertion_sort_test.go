package main

import (
	"benchForSort/additional"
	"benchForSort/mySorts"
	"testing"
)

// Бэнчмарк для 12.5.1 сортировки вставками
func BenchmarkInsertionSort(b *testing.B) {
	b.Run("small arrays insertion", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := additional.GenerateSlice(10, 10)
			b.StartTimer()
			mySorts.InsertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays insertion", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := additional.GenerateSlice(1000, 1000)
			b.StartTimer()
			mySorts.InsertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays insertion", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := additional.GenerateSlice(100000, 100000)
			b.StartTimer()
			mySorts.InsertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("small arrays insertion 12.8.4", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := additional.GenerateSlice(100, 3)
			b.StartTimer()
			mySorts.InsertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays insertion 12.8.4", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := additional.GenerateSlice(100, 300000)
			b.StartTimer()
			mySorts.InsertionSort(ar)
			b.StopTimer()
		}
	})
}
