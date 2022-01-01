package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// each meal should have chan?

	rchan := make(chan bool)
	pchan := make(chan bool)

	chorizo := meal{"chorizo", generateRandomNumber(), true}
	chopitos := meal{"chopitos", generateRandomNumber(), true}
	// fmt.Print(3)

	totalchan := make(chan int,5)
	total := chorizo.morsels + chopitos.morsels

	rchan <- chorizo.morselsAvailable
	pchan <- chopitos.morselsAvailable

	// chtotal := make(chan int)
	// chtotal <- total
	// total2 := <-chtotal

	// fmt.Println(total)
	listOfOrders := [2]meal{chorizo, chopitos}
	fmt.Println(listOfOrders)

	alice := person{"Alice"}
	bob := person{"Bob"}

	// ch := make(chan map[string]bool)

	// for {
	go alice.eat(&listOfOrders[rand.Intn(2)], &total, total)
	go bob.eat(&listOfOrders[rand.Intn(2)], &total, pchan)
	// if total == 0 {
	// 	break
	// }

	time.Sleep(time.Duration(60) * time.Minute)

	// if total != 0 {
	// 	go alice.eat(&listOfOrders[rand.Intn(2)], &total)
	// 	go bob.eat(&listOfOrders[rand.Intn(2)], &total)

	// } else {
	// 	break
	// }

	// }

}

func generateRandomNumber() int {
	return 5 + rand.Intn(6)
}

type meal struct {
	name             string
	morsels          int
	morselsAvailable bool
}

type person struct {
	name string
}

func (p person) eat(ml *meal, total *int, ch chan bool) {
	//try preventing a finished meal from being passed to func to prevent uneccesary worload

	// if ml.morsels != 0 {
	// 	ml.morselsAvailable <- true
	// 	ml.morsels--
	// 	*total--
	// 	fmt.Println(*total)
	// 	fmt.Printf("%v is enjoying some %v\n", p.name, ml.name)
	// 	timeSpentEatingMorsel := 30 + rand.Intn(271)
	// 	time.Sleep(time.Duration(timeSpentEatingMorsel) * time.Millisecond)

	// }
	// if ml.morsels != 0 {
	// 	ml.morselsAvailable <- true
	// } else {
	// 	ml.morselsAvailable <- false
	// }
	// morselAvailav := <-ml.morselsAvailable

	if ml.morsels != 0 {
		
		
		//eat
		// 	ml.morselsAvailable <- true
		// 	ml.morsels--
		// 	*total--
		// 	fmt.Println(*total)
		// 	fmt.Printf("%v is enjoying some %v\n", p.name, ml.name)
		// 	timeSpentEatingMorsel := 30 + rand.Intn(271)
		// 	time.Sleep(time.Duration(timeSpentEatingMorsel) * time.Millisecond)
	} else {
		 = false
	}

	// if ml.morselsAvailable {
	// 	ml.morselsAvailable <- true
	// 	ml.morsels--
	// 	// *total--
	// 	fmt.Println(*total)
	// 	fmt.Printf("%v is enjoying some %v\n", p.name, ml.name)
	// 	timeSpentEatingMorsel := 30 + rand.Intn(271)
	// 	time.Sleep(time.Duration(timeSpentEatingMorsel) * time.Millisecond)
	// } else {
	// 	ml.morselsAvailable <- false
	// }

	// if morselAvailav {
	// 	// ml.morselsAvailable <- true
	// 	ml.morsels--
	// 	// *total--
	// 	fmt.Println(*total)
	// 	fmt.Printf("%v is enjoying some %v\n", p.name, ml.name)
	// 	timeSpentEatingMorsel := 30 + rand.Intn(271)
	// 	time.Sleep(time.Duration(timeSpentEatingMorsel) * time.Millisecond)

	// }

}
