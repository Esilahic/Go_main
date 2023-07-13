package main

import "fmt"

func main() {
	var x int
	z := "John"
	j, k, l, _ := 2, 3, "xx", 22
	var c float64 = 45.252

	fmt.Printf("%v is of type %T\n", c, c) //printf for specificity
	fmt.Println(x, z, j, k, l)
}
