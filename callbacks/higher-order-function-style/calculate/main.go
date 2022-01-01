package main

import "fmt"

func main() {

	fmt.Println(calc(1, 2, calcProduct))

}

func calc(x, y int, operation func(x, y int) int) int {
	return operation(x, y)
}

func calcSum(x, y int) int {
	return x + y
}

func calcProduct(x, y int) int {
	return x * y
}
