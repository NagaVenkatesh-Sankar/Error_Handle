package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func invalidGetMethod() (*http.Response, error) {
	return http.Get("http://localhost:8090/sms")
}

func validPostJSON() (*http.Response, error) {
	smsContent := map[string]string{
		"to":      "987654321",
		"trom":    "998877665",
		"message": "SMS message to be sent",
	}
	reqBody, err := json.Marshal(smsContent)
	if err != nil {
		log.Println("Error occured at the JSON.")
		panic(err)
	}
	return http.Post("http://localhost:8090/sms", "application/json", bytes.NewBuffer(reqBody))
}
func invalidPostJSON() (*http.Response, error) {
	smsInvalidJSON := []byte(`{"to":"987654321","from":"998877665","message":"SMS }`)
	return http.Post("http://localhost:8090/sms", "application/json", bytes.NewBuffer(smsInvalidJSON))
}

func invalidPostDataJSON() (*http.Response, error) {
	smsInvalidContent := map[string]string{
		"name":    "NagaVenkatesh",
		"city":    "Madurai",
		"message": "SMS message to be sent",
	}
	reqBody, err := json.Marshal(smsInvalidContent)
	if err != nil {
		log.Println("Error occured at the JSON.")
		panic(err)
	}
	return http.Post("http://localhost:8090/sms", "application/json", bytes.NewBuffer(reqBody))
}

func main() {

	// 1. Invalid 'Get' method used
	// resp, err := invalidGetMethod()

	// 2. Valid JSON message payload
	// resp, err := validPostJSON()

	// 3. Invalid JSON message payload
	// resp, err := invalidPostJSON()

	// 4. Invalid JSON message payload
	resp, err := invalidPostDataJSON()

	if err != nil {
		log.Println("Error occured. Please check the connection or the Server is not reachable.")
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	//fmt.Println(resp)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	if err != nil {
		panic(err)
	}
}
