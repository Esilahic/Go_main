package main

import (
	"fmt"
)

func main() {
	es := []int{1, 2, 4, 8, 16, 32, 64, 128}
	fmt.Println(es)
	fmt.Printf("%T\n", es)

	// blank identifier to not use a returned value, this example it is the index we drop
	for _, z := range es {
		fmt.Printf("%v\n", z)

		//regular for loop
		for i := 0; i < len(es); i++ {
			fmt.Println(es[i])
		}
	}
}
