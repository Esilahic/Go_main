package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "hello"
	q := "world"
	r := "this"
	n := 3

	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Printf("%v\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", b, b)
	fmt.Printf("%v\t%T\n", c, c)
	fmt.Printf("%v\t%T\n", d, d)

	fmt.Println(*a)
	fmt.Println(*b)
	fmt.Println(*c)
	fmt.Println(*d)
}
