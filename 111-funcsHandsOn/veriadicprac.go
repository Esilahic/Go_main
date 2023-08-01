package main

import (
	"fmt"
)

func sum(ii ...int) int { // this takes individual integers and runs the for loop / requires unfurling of slice
	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}

func bar(ii []int) int { // this takes the slic of integers and runs the for loop
	t := 0
	for _, v := range ii {
		t += v
	}
	return t
}

func main() {
	x1 := []int{12, 12, 12, 12, 12, 12, 12, 12, 12}
	x := sum(x1...) // unfurling of slice to get individual integers
	fmt.Println(x)

	fmt.Println(bar(x1)) // whole slice no unfurl
}
