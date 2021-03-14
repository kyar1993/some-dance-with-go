// новое кольцо на основе предыдущего
//package main
//
//import (
//	"container/ring"
//	"fmt"
//)
//
//func main() {
//	// Создаем кольцо из пяти элементов
//	r := ring.New(5)
//	// Инициализируем элементы кольца значениями
//	// Обратите внимание мы обошли кольцо полностью один раз!
//	// Полностью!
//	for i := 0; i < r.Len(); i++ {
//		r.Value = i
//		fmt.Println("0000000000",i)
//
//		r = r.Next()
//	}
//
//	fmt.Println("0000000000",r.Value)
//
//
//	// Теперь кто-то из вас может подумать, что повторная
//	// итерация по
//	// элементам приведет к ошибке
//	// ведь мы уже в конце кольца, куда нам дальше-то двигаться?
//
//	// Из конца списка мы вновь возвращаемся автоматически в
//	// начало и
//	// перемещаемся вперед по кольцу на два элемента
//	for j := 0; j < 2; j++ {
//		r = r.Next()
//	}
//
//	fmt.Println("11111111",r.Value)
//
//	// А теперь внимание!
//	// Мы создаем новое кольцо - просто указав
//	// ссылку на следующий элемент уже существующего кольца
//	newRing := r.Next()
//	r = r.Next()
//	fmt.Printf("%d %d", newRing.Value, r.Value)
//}

// два кольца
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Создаем два кольца, каждый размером по 3 элемента
	ring1 := ring.New(3)
	ring2 := ring.New(3)

	// Получаем длины каждого из колец
	ring1Len := ring1.Len()
	ring2Len := ring2.Len()

	// Заполняем первое кольцо нулями
	for i := 0; i < ring1Len; i++ {
		fmt.Println("ring1:", i)

		ring1.Value = 0
		ring1 = ring1.Next()
	}

	fmt.Println("ring1:", ring1)

	// Заполняем второе кольцо единичками
	for j := 0; j < ring2Len; j++ {
		ring2.Value = 1
		ring2 = ring2.Next()
	}

	// Соединяем второе кольцо с первым вот таким образом
	ring1.Next().Link(ring2)
	// Так как мы соединили два кольца -
	// размер нового кольца равен сумме размеров исходных
	unionLen := ring1Len + ring2Len

	fmt.Println("unionLen:", unionLen)

	// Проверяем содержимое нового кольца
	for i := 0; i < unionLen; i++ {
		fmt.Println(ring1.Value)
		ring1 = ring1.Next()
	}
}
