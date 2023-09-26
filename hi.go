package main

import "fmt"

type Person struct {
	name string
	age  uint8
}

func main() {
	person := Person{"Anna", 42}

	fmt.Printf("Name: %s, age: %d\n", person.name, person.age)
}
