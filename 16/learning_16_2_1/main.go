// 16.2.1 Напишите программу, которая запустит n горутин,
// каждая из которых будет один раз в секунду выводить свой идентификатор в консоль.
// Идентификатором считается порядковый номер горутины. Число n задаётся
// пользователем путём ввода числа в консоль.
//
// Программа должна выполняться до тех пор, пока пользователь не нажмёт клавишу Enter.
package main

import (
	"fmt"
	"time"
)

var goroutinesCount int
var err error

func main() {
	// получаем число GoRoutines из консоли
	dataReading()

	// создаём горутины
	for i := 1; i <= goroutinesCount; i++ {
		go goroutine(i)
	}

	fmt.Println("Main finished, goroutines started\n")

	// даем время горутинам проинициализироваться и выполнить свой код
	time.Sleep(time.Second)

	// ожидаем экшена на выход
	waitingExitAction()
}

// вывод id горутины
func goroutine(id int) {
	for {
		fmt.Println("GoRoutineId: ", id)
		time.Sleep(time.Second)
	}
}

// получаем число GoRoutines из консоли
func dataReading() {
	fmt.Println("Введите количество GoRoutines:")
	_, err = fmt.Scanln(&goroutinesCount)

	if err != nil {
		fmt.Printf("При считывании количества goroutines, возникла ошибка: %s\n", err)
		fmt.Println("Допустимые значения - целые числа\n")
		dataReading()
		return
	}

	return
}

// ожидание нажатия enter - для завершения скрипта
func waitingExitAction() {
	_, err := fmt.Scanln()

	if err != nil {
		fmt.Printf("Возникла ошибка: %s\n", err)
	}
}
