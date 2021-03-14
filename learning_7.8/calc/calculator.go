package calc

import (
	"fmt"
	"os"
)

// структура калькулятор
type calculator struct{}

// Конструктор структуры калькулятор
func NewCalculator() calculator {
	return calculator{}
}

// Вычисление результата
func (calc *calculator) Calculate(firsOperand, secondOperand float64, operation string) float64 {
	var result float64

	switch operation {
	// сложение
	case "+":
		result = calc.addition(firsOperand, secondOperand)
		break
		// вычитание
	case "-":
		result = calc.subtraction(firsOperand, secondOperand)
		break
		// умножение
	case "*":
		result = calc.multiplication(firsOperand, secondOperand)
		break
		// деление
	case "/":
		result = calc.division(firsOperand, secondOperand)
		break
		// выводим ошибку - операция не поддерживается
	default:
		fmt.Println("Введённый тип операции на данный момент не поддерживается!\n")
		os.Exit(2)
	}

	return result
}

// Сложение
func (calc *calculator) addition(firstOperand, secondOperand float64) float64 {
	return firstOperand + secondOperand
}

// Вычитание
func (calc *calculator) subtraction(firstOperand, secondOperand float64) float64 {
	return firstOperand - secondOperand
}

// Умножение
func (calc *calculator) multiplication(firstOperand, secondOperand float64) float64 {
	return firstOperand * secondOperand
}

// Деление
func (calc *calculator) division(firstOperand, secondOperand float64) float64 {
	// проверить, что второй операнд != 0
	if secondOperand == 0 {
		fmt.Printf("На 0 делить нельзя!\n")
		os.Exit(3)
	}

	return firstOperand / secondOperand
}
