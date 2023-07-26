package main

import (
	"fmt"
)

func main() {
	x := 34 // 3rd
	y := 41 // 4th
	fmt.Printf("x = %v \ny = %v\n", x, y)

	switch {
	case x > 42:
		fmt.Println("x is less than 42")
	case x < 42:
		fmt.Println("x is greater than 42 ")
		fallthrough
	case x == 23:
		fmt.Println("testing fallthrough! x != 23")
	default:
		fmt.Println("this is the default case of x")
	}
}
