package main

import (
	"fmt"
)

type Person struct {
	name string
	age  uint8
}

func main() {
	person := new(Person)

	*person = Person{"Robert", 34}

	fmt.Printf("Name: %s, Age: %d\n", person.name, person.age)
}
