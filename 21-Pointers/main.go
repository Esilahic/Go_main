package main

import (
	"fmt"
)

func main() {
	x := 34
	fmt.Println(x)
	fmt.Println(&x)

	t := &x
	*t = 32
	
	fmt.Println(*t)
}
