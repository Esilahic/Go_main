package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	cs := make(chan<- int) //send only channel
	cr := make(<-chan int) //receive only channel

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cs)
	fmt.Printf("%T\n", cr)

	//general to specific converts
	// cs = c valid
	// cr = c valid

	//specific to general does not convert
	//c = cs invalid
	//c = cr invalid

}
