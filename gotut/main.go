package main

import (
	"fmt"
)

func main() {

	//EVERYTHING IN GO IS "PASSED" BY VALUE. ARGUMENTS?

	{
		// 	num := 2

		// 	var doSom = func(n int) int {
		// 		n *= 5
		// 		return n
		// 	}

		// 	result := doSom(num)

		// 	//When we pass a variable(already initialized?) into a function, we are not passing that variable at all, but merely a/the copy?/VALUE of that variable's /value
		// 	fmt.Println(result)

		// 	fmt.Print(num)
	}

	{
		// 	var dd int
		// 	//fmt.Scan also returns an error so I'm assign the retun values to two variables.
		// 	//If the value passed to "fmt.Scan" scanner input field has the same type as the type of the identifier decalred as its argument,
		// 	//the number of items scanned would be returned. If not same type e.g int required but string provided, error would occur and
		// 	//0 returned
		// 	//You can scan multiple items by using commas?
		// 	var ddd, ss = fmt.Scan(&dd)
		// 	println(ddd)
		// 	println(ss)
	}

	{
		// var scanInput string

		// fmt.Scan(&scanInput)
		// println(scanInput)
	}

	n := average(2, 4)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	total := 0.0

	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}
