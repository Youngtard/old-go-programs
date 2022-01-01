package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Hotdog.")
}

func main() {
	// var d hotdog
	http.HandleFunc("/redirecturl/", redirectURLHandler)
	http.ListenAndServe(":8080", nil)

}

func redirectURLHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Haha")
}
