package main

import (
	"fmt"
)

func foo(c chan<- int) { //send only channel
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func main() {
	c := make(chan int)

	go foo(c)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("exit")
}
