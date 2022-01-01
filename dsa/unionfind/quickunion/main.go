//LAZY APPROACH
package main

// * Pointer
// * Dereference

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	var uf UnionFind

	uf.setSimpleData()

	for {
		fmt.Println(uf.data)

		scanner.Scan()
		p, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		q, _ := strconv.Atoi(scanner.Text())
		fmt.Println(p, q)

		if uf.connected(p, q) {
			fmt.Println("CONNECTED!")
			continue
		} else {
			uf.union(p, q)

		}

		// os.Exit(0)
	}

}

type UnionFind struct {
	data []int
}

func (uf *UnionFind) root(i int) int {
	// If index != its element value, it means it has a root of a different value (it is not its own root and it has a parent which could
	// as well be its root), so find it in the while loop. While loop because the first
	// outcome/result of i in the loop may not be the root, just a parent under the root, so we have to use the parent to get to the root.
	// may also have a different root value from i

	//Goal - until we get an i/index that its id/element is its root
	for i != uf.data[i] {
		i = uf.data[i]
	}

	return i
}

func (uf *UnionFind) union(p, q int) {

	// pID := uf.data[p]
	i := uf.root(p)
	j := uf.root(q)

	uf.data[i] = uf.data[j]

}

func (uf *UnionFind) connected(p, q int) bool {
	return uf.root(p) == uf.root(q)
}

func (uf *UnionFind) setSimpleData() {

	var i int

	for ; i < 10; i++ {
		uf.data = append(uf.data, i)
	}
}
