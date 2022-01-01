package main

import (
	"fmt"
)

func main() {
	// a := "a"
	a := new(string)
	b := new(string)

	*a = "a"
	*b = "b"

	c := a
	d := "d"

	fmt.Println(a)
	fmt.Println(*a, *b, c)
	fmt.Println(&a, &b)
	fmt.Println(d)

	b = a

	fmt.Println(a, b)
	fmt.Println(&a, &b)
}
