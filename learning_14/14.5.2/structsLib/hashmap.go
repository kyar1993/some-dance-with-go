// 14.5.2. мапа для хранения кеша
package structsLib

import (
	"fmt"
	"time"
)

// это трюк для проверки типа: до тех пор пока InMemoryCache не будет
// реализовывать интерфейс Cache, программа не запустится
var _ Cache = &InMemoryCache{}

type CacheEntry struct {
	SettledAt time.Time
	Value     interface{}
}

type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
}

type InMemoryCache struct {
	expireIn time.Time              // кеш жив до...
	data     map[string]*CacheEntry // кеш мапа
}

// construct InMemoryCache
// expireIn - duration in hours
func NewInMemoryCache(expireIn time.Duration) *InMemoryCache {
	return &InMemoryCache{
		expireIn: time.Now().Add(expireIn),
		data:     make(map[string]*CacheEntry),
	}
}

// добавление значения в мапу
func (c *InMemoryCache) Set(key string, value interface{}) {
	c.data[key] = &CacheEntry{SettledAt: time.Now(), Value: value}
}

// получение значения из мапы, по ключу
func (c *InMemoryCache) Get(key string) interface{} {
	return c.data[key].Value
}

// поиск по ключу
// + вывод результатов, обработка паник
func (c *InMemoryCache) Search(key string) {
	defer errorsCatch(key)

	value := c.Get(key)

	fmt.Println("Результаты поиска:")
	fmt.Printf("Ключ %s = %s\n\n", key, value)
}

// обработка ошибок (паник)
func errorsCatch(key string) {
	if r := recover(); r != nil {
		fmt.Println("Attention! Внимание! Achtung!")
		fmt.Printf("Значение Ключа %s Не Найдено!!!\n\n", key)
	}
}
