package main

import (
	"fmt"
)

func main() {

	//functions can be assigned to variables examples below
	// An expression is a combination of values, variables, operators, and func calls that are evaluated to produce a single value.

	x := func() {
		fmt.Println("James")
	}

	y := func(s string) {
		fmt.Println("anon func name is:", s)
	}

	x()
	y("Sam")
}
