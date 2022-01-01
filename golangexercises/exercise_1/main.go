package main

import "fmt"

//Write a function which takes an integer and returns two values:
//	1)the integer divided by 2
//  2)whether or not the integer is even (true, false)
//For example
//	half(1) should return (0, false)
//	half(2) should return (1, true).

func main() {

	//Use fmt

	// println(resultOne, resultTwo)

	//Function expression exquivalent of definition at package level
	runCode := func(number int) (int, bool) {
		var isEven bool

		if number%2 == 0 {
			isEven = true
		}

		return number / 2, isEven

	}

	var resultOne, _ = runCode(44)

	println(resultOne)
	// Error compared to fmt.Println case below it which is not an Error
	// println(runCode(33))
	fmt.Println(runCode(33))
	fmt.Printf("%T", runCode)

}

// func runCode(number int) (int, bool) {
// 	var isEven bool

// 	if number%2 == 0 {
// 		isEven = true
// 	}

// 	return number / 2, isEven

// }
