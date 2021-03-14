// 14.2.2. хеш-функция на основе остатка от деления на 1000, параметр - строка
package main

import (
	"fmt"
)

func main() {
	val1 := "abc"
	val2 := "bca"

	fmt.Printf("Hash for %s = %d\n", val1, hashstr(val1))
	fmt.Printf("Hash for %s = %d\n", val2, hashstr(val2))
}

// 14.2.2. хеш-функция, параметр - строка, на основе остатка от деления на 1000.
func hashstr(text string) uint64 {
	var calc int32

	// получаем длину строки
	// используется для вычисления коэффициента
	p := int32(len(text))

	// получаем руны = int32
	for i, val := range text {
		// вычисляем коэффициент
		coef := pow(p, i)

		// вычисляем значение произведения числа на коэффициент
		calc += val * coef
	}

	// вычисляем хэш
	return uint64(calc % 1000)
}

// возведение числа в степень (рекурсивное)
// число, степень
func pow(val int32, degree int) int32 {
	// коэф 1 = числу
	if degree == 1 {
		return val

		// коэф 0 = 1
	} else if degree == 0 {
		return 1
	}

	// вычисляем значение
	return val * pow(val, degree-1)
}
