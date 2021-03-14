// структура хеш-мап (строки в качестве ключей и значений)
// методы:
// NewHashMap - Конструктор новой hash map
// Get - получение значения из hashmap, по ключу
// Add - добавление в hashmap
// Delete - получение значения из hashmap, по ключу
// PrintList - список заполненных ключей и значений
// Search - поиск
package structsLib

import "fmt"

// Объект hash map
type hashmap struct {
	list [100]string
}

// Конструктор новой hash map
func NewHashMap() *hashmap {
	return &hashmap{list: [100]string{}}
}

// добавление в hashmap
func (h *hashmap) Add(key, val string) {
	h.list[hashstr(key)] = val
}

// получение значения из hashmap, по ключу
func (h *hashmap) Get(key string) (value string, ok bool) {
	// проверить есть ли в массиве такой ключ
	keyHash := hashstr(key)

	value = h.list[keyHash]

	// значение отсутвует - ошибка
	if value == "" {
		return "", true
	}

	return value, false
}

// удаление значения по ключу
func (h *hashmap) Delete(key string) {
	h.list[hashstr(key)] = ""
}

// список заполненных ключей и значений
func (l *hashmap) PrintList() {
	for key, value := range l.list {
		if value == "" {
			continue
		}

		fmt.Println(key, value)
	}
	fmt.Println()
}

// поиск
func (l *hashmap) Search(key string) {
	val, err := l.Get(key)
	displayResults(val, err, key)
}

// отображение результатов поиска
func displayResults(value string, error bool, key string) {
	fmt.Println()

	if error {
		fmt.Printf("Не удалось найти значение по ключу: %s\n\n", key)
		return
	}

	fmt.Printf("Результаты поиска\n")
	fmt.Printf("Ключ: %s\n", key)
	fmt.Printf("Значение: %s\n", value)
}

// моя хеш-функция из 14.2.2
func hashstr(text string) uint64 {
	var calc int32

	// получаем длину строки
	// используется для вычисления коэффициента
	p := int32(len(text))

	// получаем руны = int32
	for i, val := range text {
		// вычисляем коэффициент
		coef := pow(p, i)

		// вычисляем значение произведения числа на коэффициент
		calc += val * coef
	}

	// вычисляем хэш
	return uint64(calc % 100)
}

// возведение числа в степень (рекурсивное)
// число, степень
// P.s можно было перенести в hashstr(),
// как анонимную f, решил не усложнять...Красота в простоте
func pow(val int32, degree int) int32 {
	// коэф 1 = числу
	if degree == 1 {
		return val

		// коэф 0 = 1
	} else if degree == 0 {
		return 1
	}

	// вычисляем значение
	return val * pow(val, degree-1)
}
