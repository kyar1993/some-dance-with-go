// 16.5.2 Напишите программу, которая будет каждую секунду увеличивать значение
// переменной на единицу и каждую ⅕ секунды выводить в консоль значение
// этой переменной. Программа должна выполняться n секунд и после этого закрываться.
// n вводит пользователь после запуска программы.
package main

import (
	"fmt"
	"time"
)

var timeToComplete int32
var err error
var count int

func main() {
	// получаем число GoRoutines из консоли
	dataReading()

	// задаём счётчик тиков
	// канал, который отправляет значения.
	// используем range, встроенный в канал для перебора
	// значений, поступающих каждые 200 мсек.

	// каждые 200 мс (1/5 сек) выводим значение
	ticker := time.NewTicker(200 * time.Millisecond)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at:", t)
			fmt.Println("Count:", count)
		}
	}()

	// раз в секунду увеличиваем значение переменной
	ticker2 := time.NewTicker(1 * time.Second)

	go func() {
		for t2 := range ticker2.C {
			fmt.Println("Tick at", t2)
			count++
		}
	}()

	// выставляем таймер на время введённое юзером
	time.Sleep(time.Duration(timeToComplete) * time.Second)

	// останавливаем тикер
	ticker.Stop()
	fmt.Println("Завершено!")
}

// программа должна выполняться n секунд и после этого закрываться.
func dataReading() {
	fmt.Println("Сколько времени желаете выполнять программу?")
	_, err = fmt.Scanln(&timeToComplete)

	if err != nil {
		fmt.Printf("При считывании времени на выполнение программы, возникла ошибка: %s\n", err)
		fmt.Println("Допустимые значения - целые числа\n")
		dataReading()
		return
	}

	return
}
