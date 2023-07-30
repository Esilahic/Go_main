package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func main() {
	x := doMath(1, 2, add)
	fmt.Println(x)
	y := doMath(7, 1, subtract)
	fmt.Println(y)
}
