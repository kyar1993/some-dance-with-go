// 14.3.1 тип хеш-мап (строки в качестве ключей и значений)
package main

import (
	"learning_14_3_1/structsLib"
)

func main() {
	hashmap := structsLib.NewHashMap()

	// добавляем значения
	hashmap.Add("eye-color", "green")
	hashmap.Add("size", "xxl")
	hashmap.Add("shoe-size", "45")
	hashmap.Add("name", "Peter")
	hashmap.Add("surname", "Griffin")
	hashmap.Add("age", "45")

	// при желании, можно вывести список заполненных полей
	hashmap.PrintList()

	// получение значения (поиск)
	// корректный ключ
	hashmap.Search("eye-color")

	// некорректный ключ
	hashmap.Search("aaa")

	// тестим удаление
	// выводим список
	hashmap.PrintList()
	hashmap.Delete("eye-color")
	hashmap.PrintList()
}
