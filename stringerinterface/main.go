package main

import (
	"fmt"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("the title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

func main() {
	b := book{
		title: "Forrest",
	}

	var n count = 42

	fmt.Println(b)
	fmt.Println(n)
}

// both printing func book and count bc they are using string interface.
// since int is being cast as string it is working as normal string
