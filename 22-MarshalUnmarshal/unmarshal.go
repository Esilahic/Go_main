package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int    `json:"Age"`
}

func main() {
	s := `[{"FirstName":"Lola","LastName":"Smith","Age":21},{"FirstName":"James","LastName":"Wright","Age":42}]`

	bs := []byte(s)
	fmt.Printf("%T\n%T\n", s, bs)

	People := []Person{}

	err := json.Unmarshal(bs, &People)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(People)

	for i, p := range People {
		fmt.Println("\nPerson Number:", i)
		fmt.Println(p.FirstName, p.LastName, p.Age)
	}
}
