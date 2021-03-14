package mySorts

import "math/rand"

// 12.7 алгоритм быстрой сортировки
//
// Выбираем какой-то случайный элемент из массива, который необходимо отсортировать. Этот элемент будет называться опорным.
// Перемещаем все элементы меньшие его влево, а большие, соответственно, вправо.
// Выполняем этот же алгоритм повторно для «левой» и «правой» части алгоритма.
func QuickSort(ar []int) {
	sort(ar)
}

// рекурсивная сортировка
func sort(ar []int) []int {
	// определяем длину слайса из параметров
	arrLength := len(ar)

	// слайс для элемент меньше опорного
	lessList := make([]int, 0, arrLength)

	// слайс для элемент равных опорному
	equalsList := make([]int, 0, arrLength)

	// слайс для элемент больше опорного
	moreList := make([]int, 0, arrLength)

	// в списке 1 / нет элементов - сравнивать не с чем
	if len(ar) <= 1 {
		return ar
	}

	// определяем рандомный элемент - будет опорным
	pivot := ar[rand.Intn(arrLength)]

	// перебираем список
	for _, v := range ar {
		// меньше
		if v < pivot {
			lessList = append(lessList, v)
			// больше
		} else if v > pivot {
			moreList = append(moreList, v)
			// равные опорному элементу
		} else {
			equalsList = append(equalsList, v)
		}
	}

	// ныряем в глубину списков
	// производим сортировку среза элементов меньше опорного
	lessList = sort(lessList)

	// производим сортировку среза элементов больше опорного
	moreList = sort(moreList)

	// создаём результирующий список
	var resList = make([]int, arrLength)

	// добавляем отсортированный список элементов
	// меньше + равные опорному
	resList = append(lessList, equalsList...)

	// добавляем отсортированный список элементов больше опорного
	resList = append(resList, moreList...)

	// возвращаем отсортированный подмассив / результирующий массив
	return resList
}
