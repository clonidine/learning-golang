package main

import (
	"fmt"
)

func sum(x int, y int) int {
  return x + y
}

type Person struct {
	name string
	age  uint8
}

func main() {
  a := 20
  b := 20

  result := sum(a, b)

  fmt.Printf("%d + %d = %d\n", a, b, result)
}
