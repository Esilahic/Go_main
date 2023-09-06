package main

import (
	"fmt"
)

func foo(c chan<- int) { //send only channel
	c <- 42
}

func bar(c <-chan int) { //receive only channel
	fmt.Println(<-c)
}

func main() {

	c := make(chan int)
	go foo(c)
	bar(c)

	fmt.Println("exit")
}
