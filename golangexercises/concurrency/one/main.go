package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Further challenges - TRY USING A BUFFERED CHANNEL
//replace time.Millisecond with time.Second to simulate a more realistic scenario)

func init() {
	rand.Seed(time.Now().UnixNano())
	wg.Add(1)

}

var wg sync.WaitGroup

func main() {

	alice := person{"Alice"}
	bob := person{"Bob"}

	ready := make(chan bool)

	fmt.Println("Let's go for a walk!")

	go alice.getReady(ready)
	go bob.getReady(ready)

	//main goroutine blocks here and waits for ready chan to have(receive)
	//a value that it can send to r1
	//an unbuffered channel can only perform 1 "process" at a time
	r1 := <-ready
	r2 := <-ready

	if r1 && r2 {
		close(ready)
		go armAlarm()
	}

	shoePutOn := make(chan bool)

	go alice.putOnShoes(shoePutOn)
	go bob.putOnShoes(shoePutOn)

	s1 := <-shoePutOn
	s2 := <-shoePutOn

	if s1 && s2 {
		close(shoePutOn)
		leaveHouse()
	}

	wg.Wait()

}

type person struct {
	name string
}

func (p person) getReady(ch chan bool) {

	fmt.Printf("%v started getting ready.\n", p.name)
	timeSpent := 60 + rand.Intn(31)
	time.Sleep(time.Duration(timeSpent) * time.Millisecond)
	fmt.Printf("%v spent %d seconds getting ready.\n", p.name, timeSpent)

	ch <- true

}

func armAlarm() {
	fmt.Println("Arming alarm...")
	fmt.Println("Alarm is counting down.")
	time.Sleep(time.Duration(60) * time.Millisecond)
	wg.Done()
	//Should never print to sysout before both have put on shoes.
	fmt.Println("Alarm is armed")
}

func (p person) putOnShoes(ch chan bool) {
	fmt.Printf("%v started putting on shoes.\n", p.name)
	timeSpent := 35 + rand.Intn(11)
	time.Sleep(time.Duration(timeSpent) * time.Millisecond)
	fmt.Printf("%v spent %d seconds putting on shoes\n", p.name, timeSpent)

	ch <- true
}

func leaveHouse() {
	fmt.Println("Exiting and locking door.")
}
