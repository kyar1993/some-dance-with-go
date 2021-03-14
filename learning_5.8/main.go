package main

import "test_maths/maths"
import "fmt"

func main() {
	fmt.Println("Результат сложения: " + fmt.Sprint(maths.SumResult))
	fmt.Println("Результат вычитания: " + fmt.Sprint(maths.MinusResult))
}
