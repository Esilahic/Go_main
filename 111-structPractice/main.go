package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	favIC []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		favIC: []string{"chocolate", "banana", "passion fruit with mango and guava"},
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		favIC: []string{"Strawberry", "Chocolate"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.favIC)
	fmt.Println(p2.favIC)

	for _, v := range p1.favIC {
		fmt.Println(p1.first, "fav is", v)
	}
	for _, v := range p2.favIC {
		fmt.Println(p2.first, "fav is", v)
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v)
		for _, v2 := range v.favIC {
			fmt.Println(v.first, v.last, v2)
		}
	}

	p3 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Jenny": 27,
			"Q":     87,
			"Ian":   47,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}

	for k, v := range p3.friends {
		fmt.Println(p3.first, "- friends - ", k, v)
	}

	for _, v := range p3.favDrinks {
		fmt.Println(p3.first, "- drinks - ", v)
	}
}
