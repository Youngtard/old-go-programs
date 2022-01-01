package main

import (

	//	"strings"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	num := 500
	m := map[string]int{"one": 1, "two": 2}

	mod(&num)
	mod2(m)

	// println(num)
	fmt.Printf("%v", m)
	// order := prepareDishes()
	// keys := reflect.ValueOf(order).MapKeys()
	// choice := keys[rand.Intn(5)].Interface().(string)
	// // string(choice.Interface().(string))
	// fmt.Printf("%T\n", choice)
	// // fmt.Println(choice)

}

func mod(n *int) {

	*n--
	fmt.Println(*n)
}

func mod2(m map[string]int) {
	m["two"] = 1
}

func prepareDishes() map[string]int {
	dishes := make(map[string]int)

	listOfDishes := [5]string{"chorizo",
		"chopitos",
		" pimientos de padrón",
		"croquetas",
		"patatas bravas"}

	for i := 0; i < len(listOfDishes); i++ {
		randNumber := 5 + rand.Intn(6)

		dishes[listOfDishes[i]] = randNumber

	}

	return dishes

}
