package main

import "fmt"

func main() {
	z := "sammy"
	g := 23
	x := 21.123

	fmt.Printf("%v is of type %T\n", z, z)
	fmt.Printf("%v is of type %T\n", g, g)
	fmt.Printf("%v is of type %T", x, x)
}
