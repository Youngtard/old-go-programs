package main

import "fmt"

func main() {
	button := Button{}

	// Higher order functions enables functions to be passed as callbacks
	button.addButtonEventListener("click", clicked)
	button.addButtonEventListener("hovered", hovered)

}

type Button struct {
}

func (b Button) addButtonEventListener(event string, eventHandler func()) {
	eventHandler()
}

func clicked() {
	fmt.Println("Clicked.")
}

func hovered() {
	fmt.Println("Hovered.")
}
