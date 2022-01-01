package main

import "fmt"

func main() {

	calculator := Calculator{}

	result := calculator.calc(Sum{}, 7, 6)

	fmt.Println(result)

}

type Calculate interface {
	calculate(x, y int) int
}

type Sum struct {
}

func (s Sum) calculate(x, y int) int {
	return x + y
}

type Product struct {
}

func (p Product) calculate(x, y int) int {
	return x * y
}

type Calculator struct {
}

func (c Calculator) calc(calculate Calculate, x, y int) int {
	return calculate.calculate(x, y)
}
