package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	p1 := Person{
		FirstName: "Lola",
		LastName:  "Smith",
		Age:       21,
	}

	p2 := Person{
		FirstName: "James",
		LastName:  "Wright",
		Age:       42,
	}

	people := []Person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
