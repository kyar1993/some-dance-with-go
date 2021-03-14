// Реализация Очереди Списка Вызовов Колл-Центра на основе двусвязного списка
package structsLib

import "fmt"

// ошибка при передачи несуществующего в списке индекса
var ErrWrongListIndex = fmt.Errorf("Неверный индекс списка")

// Узел списка
type CallNode struct {
	Phone string
	Next  *CallNode
	Prev  *CallNode
}

// Конструктор Нового Узла Списка
func NewCallNode(phone string) *CallNode {
	return &CallNode{phone, nil, nil}
}

// Список Вызовов
type CallsList struct {
	size int
	Head *CallNode
	Tail *CallNode
}

// Конструктор Нового Списка Вызовов
func NewCallsList() *CallsList {
	return &CallsList{0, nil, nil}
}

// получение размера списка
func (l *CallsList) Size() int {
	return l.size
}

// получение произвольного элемента списка
func (l *CallsList) Get(index int) (*CallNode, error) {
	if index < 0 || index >= l.Size() {
		return nil, ErrWrongListIndex
	}

	node := l.Head

	for i := 1; i <= index; i++ {
		node = node.Next
	}

	return node, nil
}

// обновление произвольного элемента списка
func (l *CallsList) Set(phone string, index int) error {
	if index < 0 || index >= l.Size() {
		return ErrWrongListIndex
	}
	node, err := l.Get(index)
	if err != nil {
		return err
	}
	node.Phone = phone
	return nil
}

// добавление нового элемента в конец списка
func (l *CallsList) AddToTail(phone string) {
	// список пуст - делаем вставку в начало
	if l.Head == nil {
		l.AddToHead(phone)
		return
	}

	// вставка в конец
	// создаём новый объект
	newNode := NewCallNode(phone)

	// прописываем новому предыдущий
	newNode.Prev = l.Tail

	// прописываем исходному крайнему следующий (новый)
	l.Tail.Next = newNode

	// изменяем крайний в очереди (на новый)
	l.Tail = newNode

	l.size++
}

// добавление нового элемента в конец списка
func (l *CallsList) AddToHead(phone string) {
	newNode := NewCallNode(phone)

	// список пуст - добавляем первый элемент
	// прописываем его первым и крайним
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode

		// вставляем в начало
	} else {
		// прописываем начальному - новый, как предыдущий
		l.Head.Prev = newNode

		// прописываем следующий для нового - начальный
		newNode.Next = l.Head

		// изменяем начальный в очереди (на новый)
		l.Head = newNode
	}

	l.size++
}

// вставка нового элемента в произвольную позицию
func (l *CallsList) Insert(phone string, index int) error {
	if index < 0 || index >= l.size {
		return ErrWrongListIndex
	}

	newNode := NewCallNode(phone)

	if index == 0 {
		l.AddToHead(phone)
		return nil
	}

	// получаем объект
	node, err := l.Get(index - 1)

	if err != nil {
		return err
	}

	// прописываем новому следующий - берём у элемента
	newNode.Next = node.Next

	// прописываем следующему элементу новый, как предыдущий
	newNode.Next.Prev = newNode

	// прописываем заменяемому элементу новый, как следующий
	node.Next = newNode

	// прописываем новому предыдущий - заменяемый
	newNode.Prev = node

	l.size++

	return nil
}

// удаление элемента из произвольной позиции
func (l *CallsList) Remove(index int) error {
	if index < 0 || index >= l.size {
		return ErrWrongListIndex
	}

	// получаем элемент
	node, err := l.Get(index)

	if err != nil {
		return err
	}

	// удаляем первого в очереди
	if index == 0 {
		// если единственный элемент в очереди - удаляем ссылки в списке
		if node.Next == nil {
			l.Head = nil
			l.Tail = nil
			// заменяем элемент
		} else {
			// прописываем следующий элемент - начальным в очереди
			l.Head = node.Next

			// удаляем ссылку на предыдущий элемент
			node.Next.Prev = nil
		}

		l.size--

		return nil

		// остальные
	} else {
		// заменяем ссылку на следующий элемент
		node.Prev.Next = node.Next

		// заменяем ссылку на предыдущий элемент
		node.Next.Prev = node.Prev
	}

	//node.Next = node.Next.Next

	l.size--

	return nil
}

// выводим список
func (l CallsList) Print() {
	node := l.Head
	if node != nil {
		fmt.Println("Список вызовов в ожидании ответа:")

		for node != nil {
			fmt.Printf("%s\t", node.Phone)
			node = node.Next
		}
		fmt.Printf("\n")
	} else {
		fmt.Println("Список вызовов пуст!!!")
	}
}
