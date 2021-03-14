package main

import (
	"testing"
)

// Бэнчмарк для 12.3.2 пузырьковой сортировки по-возрастанию
func BenchmarkBubbleSort(b *testing.B) {
	b.ReportAllocs()

	b.Run("small arrays bubble", func(b *testing.B) {
		b.StopTimer()
		//b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays bubble", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := generateSlice(1000, 1000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays bubble", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			ar := generateSlice(100000, 100000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})
}
