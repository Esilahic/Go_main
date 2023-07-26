package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println(("this is where the program begins"))
}

func main() {
	x := rand.Intn(250)

	fmt.Printf("x is of value %v and type %T\n", x, x)

	if x <= 100 {
		fmt.Println("x is between 0 and 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("x is between 101 and 200")
	} else {
		fmt.Println("x is between 201 and 250")
	}

	for i := 0; i < 100; i++ {
		x := rand.Intn(250)
		switch {
		case x <= 100:
			fmt.Println("between 0 and 100")
		case x >= 101 && x <= 200:
			fmt.Println("between 101 and 200")
		case x >= 201:
			fmt.Println("between 201 and 250")
		default:
			fmt.Println("not 0 - 250")
		}
	}

	for i := 0; i < 100; i++ {
		fmt.Printf("counting to 99: %v \n ", i)
	}

	for x < 200 {
		println("test")
		x++
	}
}
