package main

import (
	"fmt"
)

type number struct {
	age int
}

func (n number) speak() {
	fmt.Println("age is", n.age)
}

func main() {
	p1 := number{
		age: 15,
	}

	p2 := number{
		age: 455,
	}

	p1.speak()
	p2.speak()
}
