package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Water", "Vanilla"}
	jm := []string{"Jenny", "Money", "Water", "Chocolate"}
	fmt.Println(jb)
	fmt.Println(jm)

	xxs := [][]string{jb, jm} // holding slice of sliceString
	fmt.Println(xxs)
}
