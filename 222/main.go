package main

import (
	"math/rand"
)

func bubbleSort(ar []int) {
	// храним количество замен
	var swapCount int

	// перебираем бесконечное количество раз список значений
	for {
		// обнуляем счётчик замен
		swapCount = 0

		for j := 1; j < len(ar); j++ {
			if ar[j] < ar[j-1] {
				swapCount++
				ar[j-1], ar[j] = ar[j], ar[j-1]
			}
		}

		// останавливаем, если не было перестановок в этой итерации
		if swapCount == 0 {
			break
		}
	}
}

func generateSlice(max, size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = rand.Intn(max*2) - max
	}

	return ar
}
