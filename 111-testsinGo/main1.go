package main

import "fmt"

func main() {
	fmt.Println(Add(3, 4))
}
func Add(a int, b int) int {
	return a + b
}

// put this code into it's own file
// Test files live in the same package as the code they test.
// They are named with the following convention:
// `filename_test.go` where filename is the name
// of the file that contains the code you want to test.
