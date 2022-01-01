package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//In kasuwa code, have a otp.go file?
// If your login request is via a user supplying a username and password then a POST is preferable, as details will be sent in the HTTP messages body rather than the URL. ***Although it will still be sent plain text, unless you're encrypting via https.***
//Initialize local variables to a function before function definition or conditional statements
//api.domain.com or domain.com/api/
// Versioning API api/v1
// jwts implementations - in auth, data/payload transfer, etc
// res.Body.Close()

//session, state (management), cookies, (browser) cache, local storage

func main() {
	// secretLength := 16
	// secret := gotp.RandomSecret(secretLength)
	// fmt.Println(secret)
	// totp := gotp.NewDefaultTOTP(secret)
	// otpNow := totp.Now()
	// fmt.Println(otpNow)
	// fmt.Println(totp.Verify(otpNow, int(time.Now().Unix())))

	// /api/
	//when end point is run on browser, any output stream to browser agent is executed/writes to IO Writer
	// signingKey := []byte("")

	// jwtClaims := JWTClaims{""}

	http.HandleFunc("/api/", apiHandler)
	http.HandleFunc("/api/otp/send", sendOTPHandler)
	http.HandleFunc("/api/verify/token", provideJWTHandler)
	http.ListenAndServe(":8080", nil)

}

type otp struct {
	// Code string `json:"otp_code"` dont export? only sent through twilio API?
	SentSuccessful bool
}

type Message struct {
	PhoneNumber string
}

func sendOTPHandler(response http.ResponseWriter, request *http.Request) {
	// len := request.ContentLength
	// body := make([]byte, len)
	// b, _ := request.Body.Read(body)

	b, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		http.Error(response, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg Message
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(response, err.Error(), 500)
		return
	}

	println(msg.PhoneNumber)

	message := Message{"Test"}
	msgj, _ := json.Marshal(message)

	response.Write(msgj)
	// fmt.PrintF("%T", body)

	response.WriteHeader(http.StatusOK)

	// if otpsentthroughtwilioissuccess, respond sentsuccessful true

}

func apiHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "API says \"Do something\"")
}

type JWTClaims struct {
	// Do I need to export fields? since client is not consuming it. just use phone_number etc
	PhoneNumber string `json:"phone_number"`
	AppID       int    `json:"app_id"`
	Iat         int    `json:"iat"`
	jwt.StandardClaims
}

type JWTToken struct {
	Token string
}

// verifyTokenHandler
func provideJWTHandler(response http.ResponseWriter, request *http.Request) {

	appId := 163211
	mySigningKey := []byte("XkKSK180RRy2Hs7uvo7e1HlNsuGLNaMD")

	request.ParseForm()

	otpClaims := JWTClaims{request.PostForm["phone_number"][0], appId, int(time.Now().Unix()), jwt.StandardClaims{Subject: "sub"}}

	fmt.Println(otpClaims)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, otpClaims)

	ss, err := jwtToken.SignedString(mySigningKey)

	fmt.Println(ss, err)

	token := JWTToken{ss}

	tokenj, _ := json.Marshal(token)

	response.Write(tokenj)

	response.WriteHeader(http.StatusOK)

}
