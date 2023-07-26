package main

import (
	"fmt"
)

func main() {
	//array literal
	a := [5]int{0, 1, 2, 3, 4}
	fmt.Println(a)

	for i, v := range a {
		fmt.Printf("array is %v and of type %T\n", i, v)
	}
	fmt.Println("---------------------------------------")
	//slice hands on
	xs := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(xs)

	for j, z := range xs {
		fmt.Printf("array is %v and of type %T\n", j, z)
	}
	fmt.Println("---------------------------------------")
	fmt.Printf("xs - %v\n", xs[:5])
	fmt.Println("---------------------------------------")
	fmt.Printf("xs - %v\n", xs[5:])
	fmt.Println("---------------------------------------")
	fmt.Printf("xs - %v\n", xs[2:7])
	fmt.Println("---------------------------------------")
	fmt.Printf("xs - %v\n", xs[1:6])
	fmt.Println("---------------------------------------")

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
	fmt.Println("---------------------------------------")

	k := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(k)
	k = append(k[:3], k[6:]...)
	fmt.Printf("deleting 45, 46, and 47 - %v\n", k)
	fmt.Println("---------------------------------------")

	xi := make([]string, 0, 50)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))

	xi = append(xi, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)

	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))

	for i := 0; i < len(xi); i++ {
		fmt.Println(xi[i], i)
	}
}
