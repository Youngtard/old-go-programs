package main

import "fmt"

func main() {
	x := 1
	y := 2
	fmt.Printf("x is %d, y is %d\n", x, y)
	swap(x, y)
	fmt.Printf("x is now %d, y is now %d\n", x, y)

	var r rune = "ededed"

	fmt.Println(r)

}

func swap(x int, y int) {
	tmp := x
	x = y
	y = tmp

}
