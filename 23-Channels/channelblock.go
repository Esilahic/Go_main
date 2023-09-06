package main

import "fmt"

func main() {
	ch := make(chan int) //can also use buffered channel with 2nd arg ex. make(chan int, 1)

	go func() {
		ch <- 42 //stand alone would deadlock, but goroutine allows it to run
	}()

	fmt.Println(<-ch)
}
