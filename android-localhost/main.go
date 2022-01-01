package main

import (
	"encoding/json"
	"net/http"
)

func main() {

	http.HandleFunc("/index/", rootHandler)
	http.HandleFunc("/index/child", childHandler)
	http.HandleFunc("/index/child/infant", infantHandler)
	// http.HandleFunc("/child/", childHandler) /*Handles /child path and sub paths */
	http.ListenAndServe(":8080", nil)

}

type Message struct {
	Msg string
}

func rootHandler(response http.ResponseWriter, request *http.Request) {
	rootMessage := Message{"Hello from the index"}

	rootMessageJSON, err := json.Marshal(rootMessage)

	if err != nil {
		panic(err)
	}

	response.Write(rootMessageJSON)
}

func childHandler(response http.ResponseWriter, request *http.Request) {
	childMessage := Message{"Hello from the child"}

	userJson, err := json.Marshal(childMessage)

	if err != nil {
		panic(err)
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(userJson)

	// fmt.Fprintln(response, string(userJson))

}

func infantHandler(response http.ResponseWriter, request *http.Request) {
	infantMessage := Message{"Hello from the infant"}

	infantMessageJSON, err := json.Marshal(infantMessage)

	if err != nil {
		panic(err)
	}

	response.Write(infantMessageJSON)
}
