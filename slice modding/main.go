package main

import (
	"fmt"
)

func main() {

	// append to slice
	xi := []int{42, 43, 44, 45}
	fmt.Println(xi)
	xi = append(xi, 46, 47, 48, 49)
	fmt.Println(xi)

	//slice from slice
	xd := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xd - %v\n", xd[0:3])
	fmt.Println("---------------------------------------")
	fmt.Printf("xd - %v\n", xd[:6])
	fmt.Println("---------------------------------------")
	fmt.Printf("xd - %v\n", xd[5:])
	fmt.Println("---------------------------------------")
	fmt.Printf("xd - %v\n", xd[:])
	fmt.Println("---------------------------------------")

	// delete from a slice
	fmt.Println("deleting 5!")
	xd = append(xd[:4], xd[5:]...)
	fmt.Printf("xd - %v\n", xd)
}
