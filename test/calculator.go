/**
 * Калькулятор 6.8
 *
 * @author kyar
 */
package main

import (
	"fmt"
	"os"
)

var firstOperand, secondOperand float64
var operation string

/**
 * Вывод результа вычислений
 */
func showCalculationResult(result float64) {
	fmt.Printf("Результат вычисления: %g\n", result)
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

/**
 * Вычисление результата
 */
func calculation() float64 {
	var result float64

	switch operation {
	case "+":
		result = firstOperand + secondOperand
		break
	case "-":
		result = firstOperand - secondOperand
		break
	case "*":
		result = firstOperand * secondOperand
		break
	case "/":
		// проверить, что второй операнд != 0
		result = firstOperand / secondOperand
		break
		// выводим ошибку - операция не поддерживается
	default:
		fmt.Println("Введённый тип операции на данный момент не поддерживается!")
		os.Exit(2)
	}

	return result
}

func main() {
	//var firstOperand, secondOperand int
	//var operation string

	// ввод 1 операнда
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&firstOperand)

	// ввод типа операции
	// забирать только 1 символ далее прерывать
	fmt.Print("Введите тип операции (поддерживаемые: +,-,*,/): ")
	fmt.Scanln(&operation)

	// проверяем, что введённая операция поддерживается
	checkOperationType(operation)

	// ввод второго операнда
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&secondOperand)

	if secondOperand == 0 {
		fmt.Printf("На 0 делить нельзя!")
		os.Exit(1)
	}

	// производим расчёты
	// выводим результат
	showCalculationResult(calculation())
}
