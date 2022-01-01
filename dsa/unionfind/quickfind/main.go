package main

//EAGER APPROACH

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Be able to set arbitrary data intead of 0 - 9 for more realism
//0-9 helps though because when you arbitrarily pick a value to represent p and q indices, the element at each index is known with ease

var scanner = bufio.NewScanner(os.Stdin)

func main() {

	// uf := UnionFind{}
	var uf UnionFind

	ln := len(os.Args)

	switch ln {
	case 1:
		uf.setSimpleData(10)
	default:
		osArgsData := os.Args[1:]
		uf.setArbitraryData(osArgsData)
	}

	// os.Exit(0)

	for {
		scanner.Scan()
		p, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		q, _ := strconv.Atoi(scanner.Text())

		fmt.Println(p, q)
		// fmt.Printf("%T", p)

		if uf.connected(p, q) {
			fmt.Println("CONNECTED!")
			continue
		} else {
			uf.union(p, q)
		}

		fmt.Println(uf.data)
		fmt.Println(uf.data)

		//if count == 0, break
	}

	// fmt.Println(uf.data)

	// uf.union(0, 9)

	// fmt.Println(uf.data)

	// connected := uf.connected(0, 9)

	// fmt.Println(connected)

	// for {
	// 	scanner.Scan()
	// 	fmt.Println(scanner.Text())
	// 	// p := scanner.Text
	// }

}

type UnionFind struct {
	//Use array becasue it could be more efficient? because slices resizes dynamically and may have unecessary capacity
	data []int
}

//Union commmand: connect two objects p and q
func (uf *UnionFind) union(p, q int) {

	//Bug in code
	//E.g. If I sample p:q 0:1,1:2,*1:3, [3 3 2 3 4 5 6 7 8 9] will be returned instead of [3 3 3 3 4 5 6 7 8 9] because uf.data[p/1] would have changed value during the
	//the second iteration (of union()) to 3, so comparing uf.data[p/1] (3 - new value) to e (2) on line 75 during the third iteration would return false, so uf.data[i/2]
	//wont change to 3 i.e. if statement would not run

	//COULD TESTING THE CODE HAVE DISCOVERED THE BUG? INSTEAD OF HARD TESTING IT?

	// for i, j := range uf.data {
	// 	if uf.data[p] == j {
	// 		// uf.data[p] = uf.data[q]
	// 		uf.data[i] = uf.data[q]

	// 	}
	// }

	//SOLUTION TO ABOVE BUG, store uf.data[p] into a variable before performing union operation or TEST?
	ufP := uf.data[p]
	//Store ufQ := uf.data[q] too?

	for i, e := range uf.data {
		if ufP == e {
			// uf.data[p] = uf.data[q]
			uf.data[i] = uf.data[q]

		}
	}

	// Traditional for loop
	// for i := 0; i < len(uf.data); i++ {
	// 	if uf.data[p] == uf.data[i] {
	// 		// uf.data[p] = uf.data[q]
	// 		uf.data[i] = uf.data[q]

	// 	}
	// }

}

//Find/Connected query/Query: is there a path connnecting the two objects p and q? are they of the same component? T or F
func (uf *UnionFind) connected(p, q int) bool {

	// if uf.data[p] == uf.data[q] {
	// 	return true
	// }

	// return false

	return uf.data[p] == uf.data[q]
}

//Find number of connected components?. Component identifier for p (0 to N-1)
func (uf UnionFind) find(i int) int {

	return 1

}

//Number of components
func (uf UnionFind) count() int {
	return 1
}

//Initialize UF with 0 to n-1 integer values
//n or N
//No method overloading but optional parameters, variadic parameters, and any type <<interface{}>> parameter can be utilized
func (uf *UnionFind) setSimpleData(n int) {

	for i := 0; i < n; i++ {
		uf.data = append(uf.data, i)
	}
}

func (uf *UnionFind) setArbitraryData(osArgsData []string) {

	for _, j := range osArgsData {
		num, _ := strconv.Atoi(j)

		uf.data = append(uf.data, num)
	}

}
