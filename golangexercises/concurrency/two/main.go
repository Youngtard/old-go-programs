package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	wg.Add(5)

}

var wg sync.WaitGroup

func main() {
	// MEALS - chorizo chopitos pimientos de padrón croquetas patatas bravas
	// ch := make(chan bool, 5)
	ch := make(chan int)

	fmt.Println("Bon appétit!")
	order := prepareDishes()
	// orderList := make([]map[string]int, 5, 5)
	var totalNumberOfMorselsAvailable int
	var chorizoMorsels int
	var chopitosMorsels int
	var pimientosDePadronMorsels int
	var croquetasMorsels int
	var patatasBravas int

	for k, v := range order {
		switch k {
		case "chorizo":
			chorizoMorsels += v
		case "chopitos":
			chopitosMorsels += v
		case "pimientos de padrón":
			pimientosDePadronMorsels += v
		case "croquetas":
			croquetasMorsels += v
		case "patatas bravas":
			patatasBravas += v
		}
	}
	totalNumberOfMorselsAvailable = chorizoMorsels + chopitosMorsels + pimientosDePadronMorsels + croquetasMorsels + patatasBravas
	// ch <- totalNumberOfMorselsAvailable

	alice := person{"Alice"}
	bob := person{"Bob"}
	charlie := person{"Charlie"}
	dave := person{"Dave"}

	fmt.Printf("%v\n", order)

	// for {
	// 	go alice.eat(order, &totalNumberOfMorselsAvailable)
	// 	go bob.eat(order, &totalNumberOfMorselsAvailable)
	// 	go charlie.eat(order, &totalNumberOfMorselsAvailable)
	// 	go dave.eat(order, &totalNumberOfMorselsAvailable)

	// }

	// wg.Wait()

	// for {
	// 	if _, ok := <-ch; ok {
	// 		go alice.eat(order, &totalNumberOfMorselsAvailable, ch)
	// 		go bob.eat(order, &totalNumberOfMorselsAvailable, ch)
	// 		go charlie.eat(order, &totalNumberOfMorselsAvailable, ch)
	// 		go dave.eat(order, &totalNumberOfMorselsAvailable, ch)
	// 		// fmt.Println(totalNumberOfMorselsAvailable)
	// 		// if totalNumberOfMorselsAvailable == 0 {
	// 		// 	break
	// 		// }

	// 	} else {
	// 		break
	// 	}
	// }

	for {
		// <-ch
		// remaa := <-ch

		go alice.eat(order, &totalNumberOfMorselsAvailable, ch)
		go bob.eat(order, &totalNumberOfMorselsAvailable, ch)
		go charlie.eat(order, &totalNumberOfMorselsAvailable, ch)
		go dave.eat(order, &totalNumberOfMorselsAvailable, ch)
		remaa := <-ch
		// fmt.Println(totalNumberOfMorselsAvailable)
		if remaa == 0 {
			break
		}

		// if v, ok := <-ch {
		// 	break
		// }

	}

	// for _ = range ch {

	// 	go alice.eat(order, &totalNumberOfMorselsAvailable, ch)
	// 	go bob.eat(order, &totalNumberOfMorselsAvailable, ch)
	// 	go charlie.eat(order, &totalNumberOfMorselsAvailable, ch)
	// 	go dave.eat(order, &totalNumberOfMorselsAvailable, ch)
	// }

}

type person struct {
	name string
}

//* pass pointer to an address
// , total *int
func (p person) eat(order map[string]int, total *int, ch chan int) {
	// keys := reflect.ValueOf(order).MapKeys()
	// choice := keys[rand.Intn(5)].Interface().(string)
	//concurrent map issue
	rema := total
	listOfDishes := [5]string{"chorizo",
		"chopitos",
		"pimientos de padrón",
		"croquetas",
		"patatas bravas"}

	choice := rand.Intn(5)

	if order[listOfDishes[choice]] != 0 {
		*total--
		ch <- *rema - 1
		order[listOfDishes[choice]]--
		timeSpentEatingMorsel := 30 + rand.Intn(271)
		fmt.Printf("%v is enjoying some %v\n", p.name, listOfDishes[choice])
		time.Sleep(time.Duration(timeSpentEatingMorsel) + time.Millisecond)

	}


}

func prepareDishes() map[string]int {
	dishes := make(map[string]int)

	listOfDishes := [5]string{"chorizo",
		"chopitos",
		"pimientos de padrón",
		"croquetas",
		"patatas bravas"}

	for i := 0; i < len(listOfDishes); i++ {
		randNumber := 5 + rand.Intn(6)

		dishes[listOfDishes[i]] = randNumber

	}

	return dishes

}

eat(string meal) interface declaration

eat(string meal, ch chan) { method/func definition
	  
}

listOfDishes

bob.eat(listOfDishes[rand.Intn[randNumber]], ch)

listOfDishes

type meal struct {

	name string
	morsels int 


}



listOfDishes := []

