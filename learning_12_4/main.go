// 12.4 алгоритм сортировки выбором
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

	fmt.Println("UNSORTED LIST", numbersList)

	// 12.4.1 сортировка выбором, по-возрастанию («слева направо» поиск минимальных элементов)
	//selectionSort(numbersList)

	// 12.4.2 сортировка выбором по-убыванию («справа налево», поиск максимальных элементов)
	//selectionSortByMax(numbersList)

	// 12.4.3 сортировка выбором двунаправленная
	selectionSortBidirectional(numbersList)

	// выводим отсортированный список
	fmt.Println("\nSORTED LIST:", numbersList)
}

// 12.4.1 сортировка выбором, по-возрастанию («слева направо» поиск минимальных элементов)
func selectionSort(ar []int) {
	arrLength := len(ar)

	for i := 0; i < arrLength; i++ {
		minElementIndex := i

		for j := i + 1; j < arrLength; j++ {
			if ar[minElementIndex] > ar[j] {
				minElementIndex = j
			}
		}

		// если элемент итак на своём месте - берём следующий элемент
		if i == minElementIndex {
			continue
		}

		ar[i], ar[minElementIndex] = ar[minElementIndex], ar[i]
	}
}

// 12.4.2 сортировка выбором по-убыванию («справа налево», поиск максимальных элементов)
func selectionSortByMax(ar []int) {
	arrLength := len(ar) - 1

	for i := arrLength; i > 0; i-- {
		maxElementIndex := i

		for j := i - 1; j >= 0; j-- {
			if ar[maxElementIndex] < ar[j] {
				maxElementIndex = j
			}
		}

		// если элемент итак на своём месте - берём следующий элемент
		if i == maxElementIndex {
			continue
		}

		ar[i], ar[maxElementIndex] = ar[maxElementIndex], ar[i]
	}
}

// 12.4.3 двунаправленная сортировка выбором
// одновременный поиск максимального и минимального числа
// в конце итерации производится две операции обмена
func selectionSortBidirectional(ar []int) {
	arrLength := len(ar) - 1

	// индекс начала списка
	headIndex := 0

	// индекс конца списка
	tailIndex := arrLength

	for {
		minElementIndex := headIndex
		maxElementIndex := tailIndex

		for i := headIndex; i <= tailIndex; i++ {
			if ar[minElementIndex] > ar[i] {
				minElementIndex = i
			}

			if ar[maxElementIndex] < ar[i] {
				maxElementIndex = i
			}
		}

		// если индекс максимального эл = i, который заменили
		// заменяем индекс максимального на минимальный (после замены в нём изначальное число)
		if maxElementIndex == headIndex {
			ar[headIndex], ar[minElementIndex] = ar[minElementIndex], ar[headIndex]
			maxElementIndex = minElementIndex
			ar[tailIndex], ar[maxElementIndex] = ar[maxElementIndex], ar[tailIndex]
		} else {
			ar[headIndex], ar[minElementIndex] = ar[minElementIndex], ar[headIndex]
			ar[tailIndex], ar[maxElementIndex] = ar[maxElementIndex], ar[tailIndex]
		}

		// сдвигаем начало, конец к центру
		headIndex++
		tailIndex--

		// останавливаем, если начало и конец встретились/пересекли друг друга
		if headIndex >= tailIndex {
			break
		}
	}
}
