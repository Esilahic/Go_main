package main

import (
	"fmt"
)

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("barbar")
}

func main() {

	defer foo()
	bar()

}
