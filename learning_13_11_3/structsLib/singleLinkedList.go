// Односвязный список
package structsLib

import (
	"fmt"
)

// ошибка при передачи несуществующего в списке индекса
var ErrWrongListIndex = fmt.Errorf("Неверный индекс списка")

// Объект для списка смежных вершин
type PeakNode struct {
	Id   int
	Next *PeakNode
}

// Конструктор новой вершины
func NewPeakNode(id int) *PeakNode {
	return &PeakNode{Id: id, Next: nil}
}

// Список Вершин
type PeaksList struct {
	size int
	Head *PeakNode
}

// Конструктор новой списка вершин
func NewPeaksList() *PeaksList {
	return &PeaksList{0, nil}
}

// Размер списка
func (l *PeaksList) Size() int {
	return l.size
}

// Получение произвольного элемента списка
func (l *PeaksList) Get(index int) (*PeakNode, error) {
	if index < 0 || index >= l.Size() {
		return nil, ErrWrongListIndex
	}

	node := l.Head

	for i := 1; i <= index; i++ {
		node = node.Next
	}

	return node, nil
}

// Обновление произвольного элемента списка
func (l *PeaksList) Set(id, index int) error {
	if index < 0 || index >= l.Size() {
		return ErrWrongListIndex
	}

	node, err := l.Get(index)

	if err != nil {
		return err
	}

	//node.Weight = weight
	node.Id = id

	return nil
}

// Добавление нового элемента в конец списка
func (l *PeaksList) Add(peak *PeakNode) {
	node := l.Head

	// список пуст - выставляем head
	if node == nil {
		l.Head = peak
	} else {
		for {
			// останавливаем, если нет следующего
			if node.Next == nil {
				break
			}

			node = node.Next
		}

		// добавляем новую вершину в список
		node.Next = peak
	}

	l.size++
}

// Вставка нового элемента в произвольную позицию
func (l *PeaksList) Insert(id, index int) error {
	if index < 0 || index >= l.size {
		return ErrWrongListIndex
	}
	newNode := NewPeakNode(id)

	if index == 0 {
		l.Add(newNode)
		return nil
	}

	node, err := l.Get(index - 1)

	if err != nil {
		return err
	}

	newNode.Next = node.Next
	node.Next = newNode
	l.size++

	return nil
}

// Удаление элемента из произвольной позиции
func (l *PeaksList) Remove(index int) error {
	if index < 0 || index >= l.size {
		return ErrWrongListIndex
	}

	node, err := l.Get(index - 1)

	if err != nil {
		return err
	}

	node.Next = node.Next.Next
	l.size--

	return nil
}

// Выводим список
func (l *PeaksList) PrintList() {
	node := l.Head

	if node != nil {
		fmt.Println("Список вершин:")

		for node != nil {
			fmt.Printf("%d\t", node.Id)
			node = node.Next
		}

		fmt.Printf("\n")
	} else {
		fmt.Println("Список вершин пуст!!!")
	}
}
