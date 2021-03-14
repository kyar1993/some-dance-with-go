// 14.2.1. хеш-функция на основе остатка от деления на 1000
package main

import "fmt"

func main() {
	var val1 int64 = 123
	var val2 int64 = 12345
	var val3 int64 = 123456

	fmt.Printf("Hash for %d = %d\n", val1, hashint64(val1))
	fmt.Printf("Hash for %d = %d\n", val2, hashint64(val2))
	fmt.Printf("Hash for %d = %d\n", val3, hashint64(val3))
}

// хеш-функция на основе остатка от деления на 1000
func hashint64(val int64) uint64 {
	return uint64(val % 1000)
}
