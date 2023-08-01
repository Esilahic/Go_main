package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Println("My age is", p.age)
	fmt.Println("My name is", p.first)
}

func main() {

	p1 := person{
		first: "John",
		age:   15,
	}

	p1.speak()

}
