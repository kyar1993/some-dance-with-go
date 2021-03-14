// 14.5.2. мапа для хранения кеша
package main

import (
	"learning_14_5_2/structsLib"
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
