package main

import (
	"fmt"
)

func main() {
	//nested loop
	for i := 0; i < 5; i++ {
		fmt.Println("--")
		for j := 0; j < 5; j++ {
			fmt.Printf("outer loop: %v \n\t\t inner loop: %v\n", i, j)
		}
	}
}
