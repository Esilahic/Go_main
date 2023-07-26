package main

import (
	"fmt"
)

func foo() {
	fmt.Println("Hello")
}

func bar(s string) {
	fmt.Println("My name is", s)
}

func slam(s string) string {
	return fmt.Sprint("Yellow ", s)
}

func dad(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years:"), age
}

func main() {
	foo()

	bar("James")

	s := slam("Tea")
	fmt.Println(s)

	s1, n := dad("Dave", 54)
	fmt.Println(s1, n)
}
