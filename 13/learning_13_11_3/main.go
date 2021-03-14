// 13.11.3 Структура: ориентированный граф.
// Функции: определения кратчайшего пути между любой парой вершин.
// Рисунок во вложении.
package main

import (
	"fmt"
	"learning_13_11_3/structsLib"
)

// массив смежных узлов
// также содержит массив, но уже односвязных списков, в этот список помещаются
// указатели на вершины, фактически связанные с данной, то есть имеющие рёбра с ней
var adjacentNodesArray []*structsLib.PeaksList = make([]*structsLib.PeaksList, 0, 4)

// веса рёбер
var edgesOfWeights = map[int]map[int]int{
	1: map[int]int{
		2: 1,
		3: 7,
	},
	2: map[int]int{
		3: 4,
		4: 10,
	},
	3: map[int]int{
		4: 3,
	},
}

func main() {
	// ориентированный граф
	// заполняем массив списков смежных вершин
	adjacentNodesArrayCreate()

	length := len(adjacentNodesArray)

	// перебираем массив списков смежных вершин
	for i, list := range adjacentNodesArray {
		id := i + 1

		currentPeak := structsLib.GetPeakById(id)

		// если конечная точка - выводим результаты
		if length == id {
			showResults(currentPeak)
			break
		}

		// перебираем списки смежных вершин
		node := list.Head

		if node != nil {
			for node != nil {
				// получаем вес ребра между вершинами
				edgeWeight := edgesOfWeights[id][node.Id]

				// получаем экземпляр вершины, в котором будет уточнён вес (Релаксация)
				peak := structsLib.GetPeakById(node.Id)

				// считаем вес пути + вес ребра
				newWeight := edgeWeight + currentPeak.Weight

				// меняем вес, предыдущую вершину, если вес меньше текущего
				if newWeight < peak.Weight {
					peak.Weight = newWeight
					peak.Prev = currentPeak
				}

				node = node.Next
			}
		}
	}

}

// вывод результатов
func showResults(peak *structsLib.Peak) {
	fmt.Println("RESULT:", peak.Id, peak.Weight, peak.Prev)

	// получить путь
	prev := peak.Prev

	fmt.Printf("ROUTE: %d\t", peak.Id)

	for prev != nil {
		fmt.Printf("%d\t", prev.Id)
		prev = prev.Prev
	}

	fmt.Println()
}

// заполняем массив смежных вершин
// содержит массив, односвязных списков, в которые помещены указатели на вершины
func adjacentNodesArrayCreate() {
	secondWeight := 2
	thirdWeight := 3
	fourthWeight := 4

	// создаём списки смежных вершин
	firstList := structsLib.NewPeaksList()
	firstList.Add(structsLib.NewPeakNode(secondWeight))
	firstList.Add(structsLib.NewPeakNode(thirdWeight))

	secondList := structsLib.NewPeaksList()
	secondList.Add(structsLib.NewPeakNode(thirdWeight))
	secondList.Add(structsLib.NewPeakNode(fourthWeight))

	thirdList := structsLib.NewPeaksList()
	thirdList.Add(structsLib.NewPeakNode(fourthWeight))

	fourthList := structsLib.NewPeaksList()

	adjacentNodesArray = append(adjacentNodesArray, firstList)
	adjacentNodesArray = append(adjacentNodesArray, secondList)
	adjacentNodesArray = append(adjacentNodesArray, thirdList)
	adjacentNodesArray = append(adjacentNodesArray, fourthList)
}
