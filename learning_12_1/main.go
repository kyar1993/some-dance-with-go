package main

import "fmt"

func main() {
	//slice := []int{1, 2, 3, 4, 5}
	////slice[5] = 6 // 1
	//slice = append(slice, 5) // 2
	//slice[0] = 1 // 3

	ar := []int{1, 2, 3}
	set(ar)
	fmt.Println(ar)
}

func set(ar []int) {
	ar[1] = 4
}
