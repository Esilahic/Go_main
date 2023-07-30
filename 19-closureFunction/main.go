package main

import "fmt"

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	f := incrementor()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
