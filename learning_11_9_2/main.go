// Вы разработчик в метеослужбе. Работа идёт хорошо, у вас всё получается, и начальник, видя ваше стремление и упорство, выбрал именно вас для написания действительно важного модуля их программного комплекса.
// Вы должны написать код, который динамически рассчитывает среднюю температуру в городе от 10 метеостанций. Данные от каждой станции постоянно обновляются. Представьте, что событие изменения температуры от какой-либо метеостанции — это просто ввод соответствующей температуры через консоль и номера метеостанции, от которой она получена. При каждом изменении температуры на любой станции ваша программа должна сразу вывести среднюю температуру в городе.
package main

import (
	"fmt"
)

// список метостанций и их значений
var weatherStationsList = [10]float32{
	25.4,
	30.1,
	27.9,
	26.2,
	23.5,
	21.3,
	25.6,
	22.7,
	31.2,
	20.8,
}

const LIST_LENGTH = 10

// среднее значение
var averageValue float32

// номер метеостанции
var stationNumber int

// показания метеостанции
var stationData float32

// ошибка
var err error

// рекурсивная функция для ввода
var process func()

func main() {
	//calculateAverageValue()

	process = func() {
		// Считывание данных из командой строки
		// проверки введённых данных
		dataReading()

		// производим расчёты
		calculateAverageValue()

		// выводим результат
		showData()

		// ждём новые показания
		process()
	}

	process()
}

// отображаем среднее значение по-городу
func showData() {
	fmt.Printf("Средняя температура в городе: %.2f Град.\n\n", averageValue)
}

// расчитываем среднее значение
func calculateAverageValue() {
	var total float32

	// расчёт суммы
	for _, value := range weatherStationsList {
		total += value
	}

	averageValue = total / LIST_LENGTH
}

// Считывание данных и проверка
func dataReading() {
	// считываем номера метеостанции
	weatherStationNumberReading()

	// считываем показания метеостанции
	weatherStationDataReading()

	// изменяем значение в списке
	weatherStationsList[stationNumber-1] = stationData
}

// чтение номера метеостанции
func weatherStationNumberReading() {
	// ввод номера станции
	fmt.Print("Введите номер метеостанции: ")
	_, err = fmt.Scanln(&stationNumber)

	if err != nil {
		fmt.Printf("При номера метеостанции первого операнда, возникла ошибка: %s\n", err)
		fmt.Println("Допустимые значения - целые числа")
		weatherStationNumberReading()
	}

	// проверка номера метеостанции
	if stationNumber < 1 || stationNumber > 10 {
		fmt.Printf("Номер метеостанции %d, не поддерживается! Введите число от 1 до 10\n", stationNumber)
		weatherStationNumberReading()
	}
}

// чтение показаний метеостанции
func weatherStationDataReading() {
	// ввод номера станции
	fmt.Print("Введите показания с датчиков метеостанции: ")
	_, err = fmt.Scanln(&stationData)

	if err != nil {
		fmt.Printf("При номера метеостанции первого операнда, возникла ошибка: %s\n", err)
		fmt.Println("Допустимые значения - числа с плавающей запятой")
		weatherStationDataReading()
	}
}
