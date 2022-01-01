package main

import (
	"fmt"
	"sort"
)

func main() {

	s := []int{8, 7, 3, 1000, 4, 5, 1, 2, 3}

	greatestNumber := runCode(s...)

	fmt.Println(greatestNumber)

}

//params because if params is printed, it outputs a slice. So we name it param<s>
func runCode(params ...int) int {

	sort.Ints(params)

	return params[len(params)-1]

}
