package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("James")
	}()

	//also
	func(s string) {
		fmt.Println("anon func name is:", s)
	}("Sam")
}
