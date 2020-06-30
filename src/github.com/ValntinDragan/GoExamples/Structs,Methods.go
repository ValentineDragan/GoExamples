package main

import (
	"fmt"
)

type person struct {
	name              string
	gender            byte
	age               int
	personalityTraits []string
}

func printFullDetails(pers *person) {
	if pers.gender == 'M' {
		fmt.Printf("Mr %s, aged %d has the following personality traits: %v\n", pers.name, pers.age, pers.personalityTraits)
	} else {
		fmt.Printf("Miss %s, aged %d has the following personality traits: %v\n", pers.name, pers.age, pers.personalityTraits)
	}
}

func main() {
	people := []person{
		{name: "Valentine", gender: 'M', age: 22, personalityTraits: []string{"Friendly", "Analytical"}},
		{name: "Jina", gender: 'F', age: 22, personalityTraits: []string{"Sociable", "Hard-Working", "Greedy"}},
		{name: "Alexander", gender: 'M', age: 21, personalityTraits: []string{"Patient", "Shy"}}}

	for _, person := range people {
		printFullDetails(&person)
	}

}
