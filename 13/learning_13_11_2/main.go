// // 13.11.2. Структура: неориентированный граф.
// Функции: поиск в ширину.
// Рисунок во вложении.
package main

import (
	"fmt"
	"learning_13_11_2/structsLib"
)

// массив смежных узлов
// также содержит массив, но уже односвязных списков, в этот список помещаются
// указатели на вершины, фактически связанные с данной, то есть имеющие рёбра с ней
var adjacentNodesArray []*structsLib.PeaksList = make([]*structsLib.PeaksList, 0, 4)

// выбираем начальный элемент
var startPeakId = structsLib.FirstPeak.Id

// выбираем конечный элемент
var endPeakId = structsLib.TenthPeak.Id

// очередь обрабатываемых списков
var queue = map[int]*structsLib.PeaksList{}

// 2. Структура: неориентированный граф.
// Функции: поиск в ширину.
func main() {
	// неориентированный граф
	// заполняем массив списков смежных вершин
	adjacentNodesArrayCreate()

	// добавляем в очередь массив списков нального элемента
	queue[startPeakId] = adjacentNodesArray[startPeakId-1]

	for {
		for parentId, list := range queue {
			// если конечная точка - выводим результаты
			if parentId == endPeakId {
				showResults()
				return
			}

			// перебираем списки смежных вершин
			node := list.Head

			if node != nil {
				for node != nil {
					nodeId := node.Id

					// получаем экземпляр вершины, в котором будет уточнён вес (Релаксация)
					peak := structsLib.GetPeakById(nodeId)

					// обрабатывалась ли вершина
					// нет - добавляем список смежных элементов вершины в очередь
					// да - пропуск
					if !peak.Proccessed {
						queue[nodeId] = adjacentNodesArray[nodeId-1]
					}

					node = node.Next
				}
			}

			// отмечаем, что вершина пройдена
			parentPeak := structsLib.GetPeakById(parentId)
			parentPeak.Proccessed = true

			// удаляем список обработанной вершины из очереди
			delete(queue, parentId)
		}
	}
}

// вывод сообщения о завершении поиска
func showResults() {
	fmt.Println("Поиск успешно завершён!")
}

// заполняем массив смежных вершин
// содержит массив, односвязных списков, в которые помещены указатели на вершины
func adjacentNodesArrayCreate() {
	firstWeight := 1
	secondWeight := 2
	thirdWeight := 3
	fourthWeight := 4
	fifthWeight := 5
	sixthWeight := 6
	seventhWeight := 7
	eighthWeight := 8
	ninthWeight := 9
	tenthWeight := 10

	// создаём списки смежных вершин
	firstList := structsLib.NewPeaksList()
	firstList.Add(structsLib.NewPeakNode(secondWeight))
	firstList.Add(structsLib.NewPeakNode(thirdWeight))
	firstList.Add(structsLib.NewPeakNode(fourthWeight))

	secondList := structsLib.NewPeaksList()
	secondList.Add(structsLib.NewPeakNode(firstWeight))
	secondList.Add(structsLib.NewPeakNode(fifthWeight))

	thirdList := structsLib.NewPeaksList()
	thirdList.Add(structsLib.NewPeakNode(firstWeight))
	thirdList.Add(structsLib.NewPeakNode(sixthWeight))
	thirdList.Add(structsLib.NewPeakNode(seventhWeight))

	fourthList := structsLib.NewPeaksList()
	fourthList.Add(structsLib.NewPeakNode(firstWeight))
	fourthList.Add(structsLib.NewPeakNode(eighthWeight))

	fifthList := structsLib.NewPeaksList()
	fifthList.Add(structsLib.NewPeakNode(secondWeight))
	fifthList.Add(structsLib.NewPeakNode(ninthWeight))

	sixthList := structsLib.NewPeaksList()
	sixthList.Add(structsLib.NewPeakNode(thirdWeight))
	sixthList.Add(structsLib.NewPeakNode(tenthWeight))

	seventhList := structsLib.NewPeaksList()
	seventhList.Add(structsLib.NewPeakNode(thirdWeight))

	eighthList := structsLib.NewPeaksList()
	eighthList.Add(structsLib.NewPeakNode(fourthWeight))

	ninethList := structsLib.NewPeaksList()
	ninethList.Add(structsLib.NewPeakNode(fifthWeight))

	tenthList := structsLib.NewPeaksList()
	tenthList.Add(structsLib.NewPeakNode(sixthWeight))

	adjacentNodesArray = append(adjacentNodesArray, firstList)
	adjacentNodesArray = append(adjacentNodesArray, secondList)
	adjacentNodesArray = append(adjacentNodesArray, thirdList)
	adjacentNodesArray = append(adjacentNodesArray, fourthList)
	adjacentNodesArray = append(adjacentNodesArray, fifthList)
	adjacentNodesArray = append(adjacentNodesArray, sixthList)
	adjacentNodesArray = append(adjacentNodesArray, seventhList)
	adjacentNodesArray = append(adjacentNodesArray, eighthList)
	adjacentNodesArray = append(adjacentNodesArray, ninethList)
	adjacentNodesArray = append(adjacentNodesArray, tenthList)
}
