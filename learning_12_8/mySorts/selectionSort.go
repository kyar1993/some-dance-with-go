package mySorts

// 12.4.1 сортировка выбором, по-возрастанию («слева направо» поиск минимальных элементов)
func SelectionSort(ar []int) {
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
