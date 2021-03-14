// 16.5.1 Улучшите InMemoryCache из модуля 14, используя sync.Mutex, сделав его потокобезопасным.
//
// В качестве ответа приложите архивный файл с кодом программы из Задания 16.5.1.
package main

import (
	"learning_16_6_1/structsLib"
	"time"
)

const ttl = 2 * time.Hour

func main() {
	// создаём мапу
	inMemoryCache := structsLib.NewInMemoryCache(ttl)

	// добавляем значений
	inMemoryCache.Set("red", "pepper")
	inMemoryCache.Set("yellow", "banana")
	inMemoryCache.Set("green", "apple")

	// поиск
	// удачный поиск
	inMemoryCache.Search("yellow")

	// отсутствующий элемент
	inMemoryCache.Search("ololo")
}
