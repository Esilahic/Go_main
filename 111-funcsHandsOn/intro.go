package main

import "fmt"

func foo() int {
	return 12
}

func bar() (int, string) {
	return 83, "bar"
}

func main() {

	fmt.Println(foo())
	fmt.Println(bar())

}
