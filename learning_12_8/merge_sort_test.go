package main

import (
	"benchForSort/additional"
	"benchForSort/mySorts"
	"testing"
)

// Бэнчмарк для 12.6.1 сортировки слиянием
func BenchmarkMergeSort(b *testing.B) {
	b.Run("small arrays merge", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := additional.GenerateSlice(10, 10)
			b.StartTimer()
			mySorts.MergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays merge", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := additional.GenerateSlice(1000, 1000)
			b.StartTimer()
			mySorts.MergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays merge", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := additional.GenerateSlice(100000, 100000)
			b.StartTimer()
			mySorts.MergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("small arrays merge 12.8.4", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := additional.GenerateSlice(100, 3)
			b.StartTimer()
			mySorts.MergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays merge 12.8.4", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := additional.GenerateSlice(100, 300000)
			b.StartTimer()
			mySorts.MergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("better case merge 12.8.5", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := []int{9, 10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8}
			b.StartTimer()
			mySorts.MergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("worst case merge 12.8.5", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := []int{2, 4, 6, 8, 10, 12, 14, 16, 1, 3, 5, 7, 9, 11, 13, 15}
			b.StartTimer()
			mySorts.MergeSort(ar)
			b.StopTimer()
		}
	})
}
