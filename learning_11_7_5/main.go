// Реализация Очереди Списка Вызовов Колл-Центра на основе двусвязного списка
package main

import (
	"fmt"
	slList "learning_11_7_5/structsLib"
)

var call *slList.CallNode
var err error

// очередь
var list *slList.CallsList

func main() {
	// создаём очередь
	list = slList.NewCallsList()

	// добавляем в очередь
	list.AddToHead("+79111111111")
	list.AddToTail("+79222222222")
	list.AddToTail("+79123456789")

	// добавляем на произвольную позицию
	err = list.Insert("+79333333333", 1)

	// обрабатываем ошибку
	if err != nil {
		fmt.Println(err)
		return
	}

	// добавляем в конец очереди
	list.AddToTail("+79999999999")

	// выводим очередь
	list.Print()

	for {
		acceptNextCall()

		// выводим очередь
		list.Print()

		if list.Size() == 0 {
			// можно повесить time.Sleep(10 * time.Second)
			// чтобы проверить через времяналичие вызовов в очереди позже
			break
		}
	}
}

// принимаем следующий вызов
func acceptNextCall() {
	// получаем вызов
	call, err = list.Get(0)

	// обрабатываем ошибку
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Принят вызов от номера: %s\n\n", call.Phone)

	// удаляем обработанный вызов из очереди
	err = list.Remove(0)

	// обрабатываем ошибку
	if err != nil {
		fmt.Println(err)
		return
	}
}
