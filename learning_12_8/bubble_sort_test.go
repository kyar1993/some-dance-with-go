package main

import (
	"benchForSort/additional"
	"benchForSort/mySorts"
	"testing"
)

// Бэнчмарк для 12.3.2 пузырьковой сортировки по-возрастанию
func BenchmarkBubbleSort(b *testing.B) {
	b.Run("small arrays bubble", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := additional.GenerateSlice(10, 10)
			b.StartTimer()
			mySorts.BubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays bubble", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := additional.GenerateSlice(1000, 1000)
			b.StartTimer()
			mySorts.BubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays bubble", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := additional.GenerateSlice(100000, 100000)
			b.StartTimer()
			mySorts.BubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("small arrays bubble 12.8.4", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := additional.GenerateSlice(100, 3)
			b.StartTimer()
			mySorts.BubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays bubble 12.8.4", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := additional.GenerateSlice(100, 300000)
			b.StartTimer()
			mySorts.BubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("better case bubble 12.8.5", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := []int{1, 2, 3, 4, 5, 6}
			b.StartTimer()
			mySorts.BubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("worst case bubble 12.8.5", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := []int{6, 5, 4, 1, 2, 3}
			b.StartTimer()
			mySorts.BubbleSort(ar)
			b.StopTimer()
		}
	})
}
