package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//I could generate access_token without exchanging a code (authorization?) or credentials for it
//redirect_url had to specified in client dashboard/console and url path

func main() {

	// var context context.Context

	// http://api.genius.com/oauth/authorize?client_id=xPV3IAkzLxlOPsF1w8vS3dZW2nyC8eMVOkqZQmZwS5lKna4H3DdwxE0R2WE_lnkY&redirect_uri=http://facebook.com&scope=me&state=1234&response_type=code

	// http.HandleFunc("/", getSong)
	// http.ListenAndServe("8080", nil)
	geniusHandler()
}

func geniusHandler() {
	// baseUrl := "https://api.genius.com"
	client := http.Client{}

	const ACCESS_TOKEN string = "64dngXDjPQqqciqMbd2r2XSYi4voN-0qsiCmIOhZ8xNzfbD1h3nG1eHoW5ZG4aoR"

	// baseUrl := "https://api.genius.com"
	// searchEndpoint := "/search"
	// query parameters/args := "?q=" + placedholder_string

	// var req *http.Request
	// resp, err := http.Get("https://httpbin.org/get")
	rq, _ := http.NewRequest("GET", "https://api.genius.com/search?q=hopsin", nil)

	rq.Header.Set("Authorization", "Bearer "+ACCESS_TOKEN)

	// req.Method = "GET"
	// req.Header.Set("Authorization", "Bearer ")
	// os.Exit(1)
	// req.URL.Host = "api.genius.com"

	// req.Header.Set("Authorization", "Bearer "+ACCESS_TOKEN)

	response, err := client.Do(rq)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))

	// req.Header = make(map[string]string)
	// req.Header.Add("Authorization", "Bearer "+ACCESS_TOKEN)

	// fmt.Printf("%T\n", req)
	fmt.Printf("%T\n", rq)
	fmt.Println(rq.Header)
	os.Exit(1)

}
