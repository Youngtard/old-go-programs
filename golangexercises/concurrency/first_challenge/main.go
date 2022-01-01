package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {

	rand.Seed(time.Now().UnixNano())

}

func main() {
	//implement a range clause, or try to use channels only as first class citizens

	wg.Add(2)

	alice := Person{"Alice"}
	bob := Person{"Bob"}

	go alice.getReady()
	go bob.getReady()

	wg.Wait()

	<-armAlarm()

	go alice.putOnShoes()
	go bob.putOnShoes()

}

type Person struct {
	Name string
}

func (p Person) goForAWalk() {

}

func (p Person) getReady() {
	fmt.Println(p.Name, "started getting ready")
	randomNumber := rand.Intn(30)
	time.Sleep(time.Duration(60+randomNumber) * time.Second)
	// time.Sleep(60 + randomNumber * time.Second)

	fmt.Printf("%v spent %v seconds getting ready!\n", p.Name, 60+randomNumber)
	wg.Done()

}

func armAlarm() chan bool {
	ch := make(chan bool)
	fmt.Println("Arming Alarm!!!")
	time.Sleep(time.Duration(60))

	ch <- true

	close(ch)

	return ch
}

func (p Person) putOnShoes() {
	fmt.Println(p.Name, "started putting on shoes")
	randomNumber := rand.Intn(10)
	fmt.Printf("%v spent %v seconds putting on shoes!\n", p.Name, 35+randomNumber)
}
