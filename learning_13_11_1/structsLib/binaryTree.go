// Двоичное дерево
package structsLib

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode // левый узел
	Right *TreeNode // правый узел
}

// Конструктор нового узла дерева
func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

// Добавление нового элемента
func (t *TreeNode) Add(newNode *TreeNode) {
	value := newNode.Val

	// сравниваем родителя с новым элементом
	// новый больше - вправо
	if value > t.Val {
		// нет правого - добавляем
		if t.Right == nil {
			t.Right = newNode
			// есть - сравниваем с ним
		} else {
			t.Right.Add(newNode)
		}
		// новый меньше - влево
	} else {
		// нет правого - добавляем
		if t.Left == nil {
			t.Left = newNode
			// есть - сравниваем с ним
		} else {
			t.Left.Add(newNode)
		}
	}
}

// Поиск элемента
// принцип хранения в дереве:
// слева меньшие значения
// справа большие значения
func (t *TreeNode) Find(value int) (*TreeNode, error) {
	if t == nil {
		return nil, errors.New("Элемент не найден!")
	}

	// проверяем значение
	// совпадает - найден
	if value == t.Val {
		return t, nil
	}

	// слева меньшие значения
	if value < t.Val {
		return t.Left.Find(value)
	}

	// справа большие значения
	return t.Right.Find(value)
}

// удаление элемента
func (t *TreeNode) Remove(value int) *TreeNode {
	if t == nil {
		return nil
	}

	if value < t.Val {
		t.Left = t.Left.Remove(value)
		return t
	}

	if value > t.Val {
		t.Right = t.Right.Remove(value)
		return t
	}

	if t.Left == nil && t.Right == nil {
		return nil
	}

	if t.Left == nil {
		return t.Right
	}

	if t.Right == nil {
		return t.Left
	}

	t.Right = t.Right.Remove(t.Val)

	return t
}

// обёртка для вывода элементов дерева
func (t *TreeNode) PrintListWrapper() {
	fmt.Printf("Корень дерева:\n%d\n\n", t.Val)
	fmt.Println("Элементы дерева:")
	t.PrintList()

	fmt.Printf("\n\n")
}

// вывод элементов дерева
// сначала левую часть, затем корень и правую
func (t *TreeNode) PrintList() {
	// останавливаем рекурсию
	// если нашли лист дерева (элемент без потомков)
	if t == nil {
		return
	}

	// выводим левую часть дерева
	t.Left.PrintList()

	fmt.Printf("%d\t", t.Val)

	// выводим правую часть дерева
	t.Right.PrintList()
}

//
func (t *TreeNode) SearchResults(value int) {
	el, err := t.Find(value)

	if err != nil {
		fmt.Println("Ошибка поиска: ", err)
		return
	}

	fmt.Println("Результаты поиска:")
	fmt.Println("Значение элемента:", el.Val)

	if el.Left == nil {
		fmt.Println("Левый потомок отсутствует!")
	} else {
		fmt.Println("Значение левого потомка:", el.Left.Val)
	}

	if el.Right == nil {
		fmt.Println("Правый потомок отсутствует!")
	} else {
		fmt.Println("Значение правого потомка:", el.Right.Val)
	}
}
