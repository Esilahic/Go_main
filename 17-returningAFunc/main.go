package main

import "fmt"

//normal way so far
func bar() int {
	return 7
}

//returning a func
func foo() func() int {
	return func() int {
		return 8
	}
}

func main() {

	x := bar()
	fmt.Println(x)

	y := foo()
	fmt.Println(y())

}
