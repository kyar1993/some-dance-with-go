// Вы разработчик в метеослужбе. Работа идёт хорошо, у вас всё получается, и начальник, видя ваше стремление и упорство, выбрал именно вас для написания действительно важного модуля их программного комплекса.
// Вы должны написать код, который динамически рассчитывает среднюю температуру в городе от 10 метеостанций. Данные от каждой станции постоянно обновляются. Представьте, что событие изменения температуры от какой-либо метеостанции — это просто ввод соответствующей температуры через консоль и номера метеостанции, от которой она получена. При каждом изменении температуры на любой станции ваша программа должна сразу вывести среднюю температуру в городе.
package main

import "container/ring"

func main() {
	a := ring.IntElement{}

	//type IntElement struct {
	//	Value int
	//	Next  *IntElement
	//}
	//
	//type IntRing struct {
	//	Start *IntElement
	//	size int
	//	start *IntElement
	//}
}
