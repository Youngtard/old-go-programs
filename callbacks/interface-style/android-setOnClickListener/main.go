package main

import "fmt"

type clicked struct {
}

func main() {

	view := View{}

	view.setOnClickListener(clicked{})
}

type View struct {
}

func (v View) setOnClickListener(listener OnClickListener) {
	listener.onClick(v)
}

type OnClickListener interface {
	onClick(v View)
}

func (t clicked) onClick(v View) {
	fmt.Println("Clicked.")

}
