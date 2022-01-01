package main

import "fmt"

type person string

func main() {

	// var femi person

	// femi = "Femi"

	// femi.doing()

	btn := Button{}

	btn.addEventListener(ClickEventHandler{})
	btn.addEventListener(HoverEventHandler{})

}

// func (p person) done() {
// 	fmt.Println("Done with homework!")
// }

// func (p person) doing(handler homeworkDoneListenerHandler) {
// 	fmt.Println("Doing homework.")
// }

type ButtonEventListener interface {
	// Event Handling. Callback Method
	handleEvent()
}

type ClickEventHandler struct {
}

type HoverEventHandler struct {
}

// Implementation of callback function declared at the EventListener interface
func (handler ClickEventHandler) handleEvent() {
	fmt.Println("Clicked.")
}

// Callback method
func (handler HoverEventHandler) handleEvent() {
	fmt.Println("Hovered.")
}

type Button struct {
}

// Caller Function from Button that calls back handleEvent()
func (btn Button) addEventListener(listener ButtonEventListener) {
	// Callback Function (handleEvent()) - invoked based on type that implemented the interface
	listener.handleEvent()
}
