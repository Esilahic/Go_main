package main

import (
	"fmt"
)

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("My name is", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("secret agent is", sa.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}

	p2 := person{
		first: "Ken",
	}

	saySomething(sa1)
	saySomething(p2)

	//sa1.speak()
	//p2.speak()
}
