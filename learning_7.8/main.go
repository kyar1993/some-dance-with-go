/*
 * Модуль калькулятор
 * @todo вынести взаимодействия с юзером в отдельный пакет, к сожалению не соответствует ТЗ
 */
package main

import (
	"fmt"
	"learning_7.8/calc"
	"os"
)

var err error
var firstOperand, secondOperand, result float64
var operation string
var process func()

func main() {
	process = func() {
		fmt.Println("start")

		// Считывание данных из командой строки
		// проверки введённых данных
		dataReading()

		// создаём экземпляр структуры
		calcStruct := calc.NewCalculator()

		// производим расчёты
		result = calcStruct.Calculate(
			firstOperand,
			secondOperand,
			operation,
		)

		// выводим результат
		showCalculationResult(result)

		// посчитать что-то ещё?
		needToRepeat()
	}

	process()
}

// взаимодействие с пользователем: иные расчёты?
func needToRepeat() {
	var repeat string

	fmt.Println("Посчитать за Вас что-то ещё? Y/n")
	_, err = fmt.Scanln(&repeat)

	// yes
	if repeat == "Y" || repeat == "Д" || repeat == "" {
		process()
		return
	}

	// no or some unhandled error
	fmt.Println("Возвращайтесь за новыми расчётами!!!")
}

// Считывание данных из командой строки
// проверки введённых данных
func dataReading() {
	// ввод 1 операнда
	fmt.Print("Введите первое число: ")
	_, err = fmt.Scanln(&firstOperand)

	if err != nil {
		fmt.Printf("При сканировании первого операнда, возникла ошибка: %s\n", err)
		fmt.Println("Допустимые значения - целые числа / числа с плавающей точкой")
		os.Exit(1)
	}

	// ввод типа операции
	// забирать только 1 символ далее прерывать
	fmt.Print("Введите тип операции (поддерживаемые: +,-,*,/): ")
	fmt.Scanln(&operation)

	// проверяем, что введённая операция поддерживается
	checkOperationType(operation)

	// ввод второго операнда
	fmt.Print("Введите второе число: ")
	_, err = fmt.Scanln(&secondOperand)

	if err != nil {
		fmt.Printf("При сканировании второго операнда, возникла ошибка: %s\n", err)
		fmt.Println("Допустимые значения - целые числа / числа с плавающей точкой")
		os.Exit(1)
	}

	if secondOperand == 0 {
		fmt.Println("На 0 делить нельзя!")
		os.Exit(1)
	}
}

/**
 * Проверка поддержки операции вычисления
 */
func checkOperationType(operation string) {
	acceptedOperations := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}

	if !acceptedOperations[operation] {
		// прерываем программу
		fmt.Printf("Введённый тип операции %s, не поддерживается! Введите одно из следующих действий: + - * /\n", operation)
		os.Exit(1)
	}
}

// Вывод результа вычислений
func showCalculationResult(result float64) {
	fmt.Printf("Результат вычисления: %g\n\n", result)
}
