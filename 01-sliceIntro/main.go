package main

import (
	"fmt"
)

func main() {
	//slices
	s := []string{"hello", "world"}
	fmt.Println(s)

	var es = []int{1, 2, 4, 8, 16, 32, 64, 128}
	fmt.Println(len(es))
	for i, v := range es {
		fmt.Printf("2**%d = %d\n", i, v)
	}

}
