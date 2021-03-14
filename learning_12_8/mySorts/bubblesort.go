package mySorts

// 12.3.2 пузырьковая сортировка по-возрастанию
func BubbleSort(ar []int) {
	//arLen := len(ar)
	//for i := 0; true; i++ {
	//	sorted := true
	//	for i := 1; i < arLen; i++ {
	//		if ar[i-1] > ar[i] {
	//			ar[i-1], ar[i] = ar[i], ar[i-1]
	//			sorted = false
	//		}
	//	}
	//	if sorted {
	//		return
	//	}
	//	arLen--
	//}

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
			//break
			return
		}
	}
}
