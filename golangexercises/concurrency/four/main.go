package main

import (

	// "fmt"
	"fmt"
	"math/rand"
)

func main() {

	chorizo := meal{"chorizo", generateRandomNumber(), true}
	chopitos := meal{"chopitos", generateRandomNumber(), true}
	listOfOrders := [2]meal{chorizo, chopitos}
	alice := person{"Alice"}
	bob := person{"Bob"}

	for {
		go alice.eat(&listOfOrders[rand.Intn(2)])
		go bob.eat(&listOfOrders[rand.Intn(2)])

		if chorizo.morsels == 0 && chopitos.morsels == 0 {
			break
		}

	}

}

func generateRandomNumber() int {
	return 5 + rand.Intn(6)
}

type meal struct {
	name      string
	morsels   int
	availbale bool
}

type person struct {
	name string
}

func (p person) eat(ml *meal) {
	if ml.availbale {
		ml.morsels--
		fmt.Printf("%v is enjoying some %v\n", p.name, ml.name)

	}
	if ml.morsels == 0 {
		ml.availbale = false
	}

	for {

	}

}
