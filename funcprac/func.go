package main

import "fmt"

// only function main is compiled, but outside of main functions can be called.

/* example syntax

func function_name( [parameter list (can be empty) ] ) [return_types]
{
   body of the function
}

*/

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(44, 44))
}
