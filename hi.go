package main

import (
	"fmt"
)

type Person struct {
	name string
	age  uint8
}

func main() {
	persons := []Person{{"Anna", 42}, {"Robert", 50}}

	for i := 0; i < len(persons); i++ {
		person := persons[i]

		fmt.Printf("Name: %s, Age: %d\n", person.name, person.age)
	}
}
