package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"hotpTest/crypto"
	"io/ioutil"
	"net/http"
	"os"
)

// A simple program to test a hotp based api
// TODO: Add functionality to set values directly from enviroment variables
func main() {
	var passLength int
	var t0, interval int64
	var encryption, secretKey, targetApi, userid, message, passLenFormat string

	// TODO: Improve parameter parsing
	// Begin parsing
	flag.StringVar(&secretKey, "secret", "defaultPass", "The secret to encrpt hotp")
	flag.StringVar(&encryption, "encryption", "sha512", "Type of encryption, default is sha512")
	flag.StringVar(&targetApi, "target", "", "The site to post the message using hotp.")
	flag.StringVar(&userid, "userid", "default@gmail.com", "Your user id")
	flag.StringVar(&message, "message", "{userid: abc, lol: \"sample message.\"}", "Message to send")

	flag.Int64Var(&t0, "initial", 0, "t0")
	flag.Int64Var(&interval, "interval", 30, "Interval between new hotp")

	flag.IntVar(&passLength, "lenght", 10, "Length of the hotp.")

	flag.Parse()
	// End parsing

	// Format for padding with leading zeroes if password length is less than passLength
	passLenFormat = fmt.Sprintf("%%0%dd", passLength)

	password := hotp.CalcHotp(encryption, []byte(secretKey), t0, interval, passLength)

	auth := base64.StdEncoding.EncodeToString([]byte(userid + ":" + fmt.Sprintf(passLenFormat, password)))

	fmt.Printf("%s\n", fmt.Sprintf(passLenFormat, password))
	fmt.Printf("auth: %s\n", auth)

	if targetApi == "" {
		fmt.Println("Skipping Api Test")
		os.Exit(0)
	}

	fmt.Printf("Initiating api test. Sending http request to %s\n", targetApi)
	req, _ := http.NewRequest("POST", targetApi, bytes.NewBuffer([]byte(message)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+string(auth))
	req.Header.Set("Accept", "*/*")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status: ", response.Status)
	fmt.Println("Headers: ", response.Header)
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Body: ", string(body))
}
