// 12.3 алгоритм пузырьковой сортировка
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	// необходимо для того, чтобы рандом был похож на рандомный
	rand.Seed(time.Now().UnixNano())
}

func main() {
	numbersList := make([]int, 50)

	for i := range numbersList {
		numbersList[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}

	// выводим неотсортированный список
	fmt.Println("UNSORTED:", numbersList)

	// 12.3.3 сортируем список по-возрастанию
	//bubbleSort(numbersList)

	// 12.3.4 сортируем список по-убыванию
	//bubbleSortReversed(numbersList)

	// 12.3.5 сортируем список по-возрастанию рекурсивно
	bubbleSortRecursive(numbersList)

	// выводим отсортированный список
	fmt.Println("\nSORTED LIST:", numbersList)
}

// 12.3.2 пузырьковая сортировка по-возрастанию
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

// 12.3.3 пузырьковая сортировка по-возрастанию
//
// оптимизированная версия 12.3.2
// (сортировка прерываетс, если не было перестановок в этой итерации)
func bubbleSortOptimised(ar []int) {
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

// 12.3.4 пузырьковая сортировка по-возрастанию
func bubbleSortReversed(ar []int) {
	// храним количество замен
	var swapCount int

	// перебираем бесконечное количество раз список значений
	for {
		// обнуляем счётчик замен
		swapCount = 0

		for j := 1; j < len(ar); j++ {
			if ar[j] > ar[j-1] {
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

// 12.3.5 пузырьковая сортировка по-возрастанию, на основе рекурсивной функции
func bubbleSortRecursive(ar []int) {
	// храним количество замен
	// обнуляем счётчик
	swapCount := 0

	// P.s для DRY, можно вынести этот цикл в отдельную функцию
	for j := 1; j < len(ar); j++ {
		if ar[j] < ar[j-1] {
			swapCount++
			ar[j-1], ar[j] = ar[j], ar[j-1]
		}
	}

	// если были перестановки в этой итерации - вызываем повторную сортировку
	if swapCount != 0 {
		bubbleSortRecursive(ar)
	}
}
