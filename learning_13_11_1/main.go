// 13.11.1. Структура: двоичное дерево.
// Функции: вставка элемента, удаление элемента, поиск элемента, печать элементов дерева.
package main

import (
	"learning_13_11_1/structsLib"
)

// 1. Структура: двоичное дерево.
// Функции: вставка элемента, удаление элемента, поиск элемента, печать элементов дерева.
func main() {
	// создаём корень дерева
	treeRoot := structsLib.NewTreeNode(10)

	// вставка элементов
	treeRoot.Add(structsLib.NewTreeNode(17))
	treeRoot.Add(structsLib.NewTreeNode(4))
	treeRoot.Add(structsLib.NewTreeNode(5))
	treeRoot.Add(structsLib.NewTreeNode(6))
	treeRoot.Add(structsLib.NewTreeNode(3))

	// выводим список элементов дерева
	treeRoot.PrintListWrapper()

	// корректный поиск элемента
	//treeRoot.SearchResults(5)

	// некорректный поиск элемента
	//treeRoot.SearchResults(200)

	// удаляем элемент
	treeRoot.Remove(5)

	// выводим список элементов дерева
	treeRoot.PrintListWrapper()
}
