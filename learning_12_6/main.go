// 12.6 алгоритм сортировки слиянием
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

	fmt.Println("UNSORTED LIST:", numbersList)

	// 12.6.1 сортировка слиянием
	sortedList := mergeSort(numbersList)

	// выводим отсортированный список
	fmt.Println("\nSORTED LIST:", sortedList)
}

// 12.6.1 сортировка слиянием
func mergeSort(ar []int) []int {
	// получаем длину слайса
	length := len(ar)

	// если больше 1 элемента в списке - сортируем
	if length >= 2 {
		// получаем середину слайса
		middle := length / 2

		// сортируем слайсы
		// мерджим отсортированные слайсы
		ar = merge(
			mergeSort(ar[:middle]),
			mergeSort(ar[middle:]),
		)
	}

	return ar
}

// выполняем сравнения
func merge(left []int, right []int) []int {
	resArr := make([]int, 0)

	for len(left) > 0 && len(right) > 0 {
		// сравниваем элементы
		if left[0] < right[0] {
			// левый больше - записываем в результирующий слайс
			resArr = append(resArr, left[0])

			// на следующем шаге будет список с 1 элемента, до конца
			left = left[1:]

			// правый больше - записываем в результирующий слайс
		} else {
			resArr = append(resArr, right[0])

			// на следующем шаге будет список с 1 элемента, до конца
			right = right[1:]
		}
	}

	// если остались элементы в слайсах - добавляем
	if len(left) > 0 {
		resArr = append(resArr, left...)
	}
	// если остались элементы в слайсах - добавляем
	if len(right) > 0 {
		resArr = append(resArr, right...)
	}

	return resArr
}
