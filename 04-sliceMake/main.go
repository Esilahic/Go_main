package main

import (
	"fmt"
)

func main() {
	xi := make([]int, 0, 10) // holding no current value, but capacity of 10 available
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	// appening to made slice bc of scope once append is called then it would be filled
	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("xi - %v\n", xi)
	fmt.Println("--------------------------")
}
