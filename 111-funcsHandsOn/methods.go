package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Println()
	fmt.Println("My name is", p.first, "and my age is", p.age)
}

func main() {

	p1 := person{
		first: "John",
		age:   15,
	}

	p1.speak()

}
