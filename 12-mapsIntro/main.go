package main

import (
	"fmt"
)

func main() {
	mp := map[string]int{
		"James A.": 66,
		"James B.": 22,
		"James C.": 12,
	}

	delete(mp, "James A.")

	for z, v := range mp {
		fmt.Printf("%s - %d \n", z, v)
	}

	// v, ok := mp["James C."]
	// if ok {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println("no go")
	// }

	// best practice
	if v, ok := mp["James A."]; !ok {
		fmt.Println("no go")
	} else {
		fmt.Println("the value is", v)
	}
}
