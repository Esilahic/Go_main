package main

import "fmt"

func main() {
	x, y, z := 747, 911, 90210

	fmt.Printf("%d \t %b \t %#x\n", x, x, x)
	fmt.Printf("%d \t %b \t %#x\n", y, y, y)
	fmt.Printf("%d \t %b \t %#x", z, z, z)
}
