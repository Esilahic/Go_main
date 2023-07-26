package main

import "fmt"

func sum(ii ...int) int {
	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}

func main() {
	x1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(x1...) // unfurls slice into individual values
	fmt.Println(x)
}
