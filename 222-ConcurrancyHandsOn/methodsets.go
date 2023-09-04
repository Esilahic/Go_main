package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("I am", p.first, p.last)
}

func saySomething(h human) {
	h.speak()
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
	}

	// saySomething(p1) // This will not work
	saySomething(&p1)

}
