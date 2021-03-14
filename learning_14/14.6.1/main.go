// 14.6.1 программа, которая находит общие элементы в двух массивах
package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var messagesMap = map[int]string{
	0: "",
	1: "Введите размер первого массива:",
	2: "Введите размер второго массива:",
	3: "При сканировании операнда, возникла ошибка: %s\n",
	4: "Размер массива должен быть числом, больше 0!\n",
	5: "Введите первый массив",
	6: "Введите второй массив",
	7: "Общих элементов не найдено",
}

// переменные размеров мап
var firstArraySize, secondArraySize int

// временная переменая
var temp string

// мапы
var firstMap, secondMap map[string]string

// слайс общих элементов
var splitSlice = make([]string, 0, 10)

func main() {
	// считываем значения
	dataReading()

	// сравниваем
	for _, v := range firstMap {
		if secondMap[v] != "" {
			splitSlice = append(splitSlice, v)
		}
	}

	// общих элементов не найдено
	if len(splitSlice) == 0 {
		fmt.Println()
		fmt.Println(messagesMap[7])
		return
	}

	// выводим общие элементы
	fmt.Println()
	fmt.Println("Список общих элементов:")
	fmt.Println(splitSlice)
}

// считывание данных из командой строки
// проверки введённых данных
func dataReading() {
	// считываем размеры массивов
	firstMapCountReading()
	secondMapCountReading()

	fmt.Println()
	fmt.Println("РАЗМЕРЫ МАССИВОВ:")
	fmt.Println("Первый:", firstArraySize)
	fmt.Println("Второй:", secondArraySize)
	fmt.Println()

	// считываем значения массивов
	fmt.Println(messagesMap[5])
	firstMapValuesReading()
	fmt.Println()

	fmt.Println(messagesMap[6])
	secondMapValuesReading()
	fmt.Println()

	fmt.Println("МАССИВЫ:")
	fmt.Println("Первый:", firstMap)
	fmt.Println("Второй:", secondMap)
}

// проверка корректности введённого размера массивов
func mapCountChecker(saveTo *string) bool {
	val := *saveTo

	convertedVal, err := strconv.ParseInt(val, 10, 64)

	if err != nil || convertedVal == 0 {
		fmt.Println(messagesMap[4])
		return false
	}

	return true
}

// считывание размера первого массива
func firstMapCountReading() {
	temp = readingBase(&temp, 1, mapCountChecker)

	convertedVal, _ := strconv.Atoi(temp)

	firstArraySize = convertedVal

	// создать мапу
	firstMap = make(map[string]string, convertedVal)

	// чистим временную переменную
	tempClean()
}

// считывание размера второго массива
func secondMapCountReading() {
	temp = readingBase(&temp, 2, mapCountChecker)

	convertedVal, _ := strconv.Atoi(temp)

	secondArraySize = convertedVal

	// создать мапу
	secondMap = make(map[string]string, convertedVal)

	// чистим временную переменную
	tempClean()
}

// базовый метод чтения данных из консоли
// saveTo - в какую переменную сохраняем, считанное значение
func readingBase(saveTo *string, headerMsg int, callback func(*string) bool) string {
	// отображение заголовка
	if headerMsg != 0 {
		fmt.Println(messagesMap[headerMsg])
	}

	_, err := fmt.Scanln(saveTo)

	if err != nil {
		fmt.Printf(messagesMap[3], err)
		os.Exit(1)
	}

	// введено некорректное значение
	if !callback(saveTo) {
		return readingBase(saveTo, headerMsg, callback)
	}

	return *saveTo
}

// считывание данных для первой мапы
func firstMapValuesReading() {
	for i := 0; i < firstArraySize; i++ {
		temp = readingBase(&temp, 0, valuesChecker)

		// пишем в мапу
		firstMap[temp] = temp

		// очищаем временную переменную
		tempClean()
	}
}

// считывание данных для второй мапы
func secondMapValuesReading() {
	for i := 0; i < secondArraySize; i++ {
		temp = readingBase(&temp, 0, valuesChecker)

		// пишем в мапу
		secondMap[temp] = temp

		// очищаем временную переменную
		tempClean()
	}
}

// проверка корректности размера массивов
func valuesChecker(saveTo *string) bool {
	val := *saveTo

	match, _ := regexp.MatchString("[0-9]", val)

	if !match {
		return false
	}

	return true
}

// очистка временной переменной
func tempClean() {
	temp = ""
}
