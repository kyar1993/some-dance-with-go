/*
 * Модуль 10.7
 */
package main

import "fmt"

func main() {
	// слайс
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// слайс с отрицательными числами
	arrWithNegatives := []int{1, 10, 2, -3, 4, 5, 2, 6, -7, 2, 8, 9, 10}

	//mapWithNegatives := map[int]int{
	//	0:  1,
	//	1:  10,
	//	3:  2,
	//	4:  -3,
	//	5:  4,
	//	6:  5,
	//	7:  2,
	//	8:  6,
	//	9:  -7,
	//	10: 1,
	//	11: 8,
	//	12: 9,
	//	13: 10,
	//}

	// поиск максимального
	//res, err := findMax(arr)

	// поиск наименьшего отрицательного
	//res, err := findMaxNegative(arrWithNegatives)

	// поиск наиболее повторяющегося
	//res, err := findMostOftenRepeated(arrWithNegatives)

	// 10.7.3 поиск наиболее повторяющегося с помощью мап
	//res, err := findMostOftenRepeatedWithMap(mapWithNegatives)

	// 10.7.4 удаляет все отрицательные числа
	//res, err := trimNegative(arrWithNegatives)

	// 10.7.5 числа меньше среднего арифметического
	res, err := lessThanAverage(arrWithNegatives)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", res)
	}
}

// среднее арифметическое
func sliceAverage(xs []int) int {
	total := 0

	for _, v := range xs {
		total += v
	}

	return total / len(xs)
}

// 10.7.5 числа меньше среднего арифметического
func lessThanAverage(list []int) (res []int, err error) {
	average := sliceAverage(list)

	for _, number := range list {
		if number < average {
			res = append(res, number)
		}
	}

	return res, nil
}

// 10.7.4 удаляет все отрицательные элементы
func trimNegative(list []int) (res []int, err error) {
	for _, number := range list {
		if number >= 0 {
			res = append(res, number)
		}
	}

	return res, nil
}

// 10.7.3 Наиболее часто повторяющиеся на основе map

// не доделано, т.к чтобы было O(n) нужно юзать 1 цикл
// в этом случае создадим мапу со всеми элементами и их количеством повторений
// в итоге придётся добавлять фильтр (перебирать), чтобы исключить менее повторяющиеся
func findMostOftenRepeatedWithMap(list map[int]int) (resultMap map[int]int, err error) {
	//if len(list) == 0 {
	//	return 0, fmt.Errorf("could not found repeated numbers in empty slice")
	//}

	// создаём новую мапу - для возврата
	result := make(map[int]int)

	// перебираем мапу
	for number := range list {
		count := result[number]
		count++

		result[number] = count

		fmt.Println("!!!!!!!!!!: ", result)
	}

	return result, nil
}

// Наиболее часто повторяющиеся на основе array
func findMostOftenRepeated(array []int) (mostOften int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found repeated numbers in empty slice")
	}

	var maxNumber, maxCount = 0, 0
	for i, number := range array {
		currentCount := 0
		for _, numberToCompare := range array[i:] {
			if number == numberToCompare {
				currentCount++
			}
		}

		if currentCount > maxCount {
			maxNumber = i
			maxCount = currentCount
		}
	}

	fmt.Println("111111111: ", maxNumber)
	fmt.Println("222222222: ", maxCount)

	return maxNumber, nil
}

// максимальное отрицательное число (а в случае отсутствия отрицательных чисел — ошибку).
func findMaxNegative(array []int) (max int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found max in empty slice")
	}

	maxNegative := array[0]
	for _, val := range array[1:] {
		if val < maxNegative {
			maxNegative = val
		}
	}

	if maxNegative >= 0 {
		return 0, fmt.Errorf("Error: Could not found negative numbers in slice")
	}

	return maxNegative, nil
}

func findMax(array []int) (max int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found max in empty slice")
	}

	max = array[0]
	for _, val := range array[1:] {
		if val > max {
			max = val
		}
	}

	return max, nil
}

// получаем список подозреваемых из "БД"
//func fetchSuspectsList() map[string]Man {
//	return map[string]Man{
//		"Rick":   Man{Name: "Rick", LastName: "Sanchez", Age: 70, Gender: manGender, Crimes: 323},
//		"Morty":  Man{Name: "Morty", LastName: "Smith", Age: 14, Gender: manGender, Crimes: 23},
//		"Peter":  Man{Name: "Peter", LastName: "Griffin", Age: 45, Gender: manGender, Crimes: 54},
//		"Brian":  Man{Name: "Brian", LastName: "Griffin", Age: 6, Gender: manGender, Crimes: 14},
//		"Stewie": Man{Name: "Stewie", LastName: "Griffin", Age: 2, Gender: manGender, Crimes: 127},
//		"Homer":  Man{Name: "Homer", LastName: "Simpson", Age: 38, Gender: manGender, Crimes: 100500},
//		"Marge":  Man{Name: "Marge", LastName: "Simpson", Age: 34, Gender: womanGender, Crimes: 4},
//		"Bart":   Man{Name: "Bart", LastName: "Simpson", Age: 10, Gender: manGender, Crimes: 219},
//		"Lisa":   Man{Name: "Lisa", LastName: "Simpson", Age: 8, Gender: womanGender, Crimes: 9},
//		"Maggie": Man{Name: "Maggie", LastName: "Simpson", Age: 1, Gender: womanGender, Crimes: 1},
//	}
//}
