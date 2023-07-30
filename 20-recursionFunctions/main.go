package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factorLoop(x int) int {
	total := 1
	for x > 0 {
		total *= x
		x--
	}
	return total
}

func main() {

	fmt.Println(factorial(4))
	fmt.Println(factorLoop(4))
}
