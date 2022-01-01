package main

import (
	"fmt"
	"time"
)

func main() {
	// If Javascript, 2 would be printed first because "JavaScript is an event driven language (purpose of Rx Java?). This means that instead of waiting for a response before moving on, JavaScript will keep executing while listening for other events"

	first()
	// go first()
	second()

	// Cannot work in java because java doesn't have functions but methods? so use interface style
	doHomeWork(done)

}

func first() {
	time.Sleep(2 * time.Second)
	fmt.Println(1)
}

func second() {
	fmt.Println(2)
}

func doHomeWork(doneCallback func()) {
	fmt.Println("Doing homework")
	doneCallback()
}

func done() {
	fmt.Println("Done with homework")
}

func clickEventHandler() {

}

// package main

// import "fmt"

// type person string

// func main() {

// 	// var femi person

// 	// femi = "Femi"

// 	// femi.doing()

// 	btn := Button{}

// 	btn.onClick(ClickHandler{})

// }

// // func (p person) done() {
// // 	fmt.Println("Done with homework!")
// // }

// // func (p person) doing(handler homeworkDoneListenerHandler) {
// // 	fmt.Println("Doing homework.")
// // }

// type ClickEventHandler interface {
// 	handleClick()
// }

// type ClickHandler struct {
// }

// func (ch ClickHandler) handleClick() {
// 	fmt.Println("Clicked.")
// }

// type Button struct {
// }

// func (btn Button) onClick(clickHandler ClickEventHandler) {
// 	clickHandler.handleClick()

// }
