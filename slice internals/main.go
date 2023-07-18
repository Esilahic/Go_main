package main

import "fmt"

func main() {
	// using make and copy to change strings that start out same, but keep changes to only one if modified
	a := []int{0, 1, 2, 3, 4, 5, 6}
	b := make([]int, 7)
	copy(b, a)
	fmt.Println("a", a)
	fmt.Println("b", b)

	a[0] = 99
	fmt.Println("new a", a)
	fmt.Println("same b", b)
}
