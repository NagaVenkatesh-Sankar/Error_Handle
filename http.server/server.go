package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

type smsSchema struct {
	From    string `json:"from" validate:"required"`
	To      string `json:"to" validate:"required"`
	Message string `json:"message"`
}

func sms(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		return
	}

	// Read request body.
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Body read error, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		fmt.Fprintf(w, err.Error())
		return
	}

	// Parse body as json.
	var smsContext smsSchema
	if err = json.Unmarshal(body, &smsContext); err != nil {
		log.Printf("Body parse error, %v", err)
		w.WriteHeader(400) // Return 400 Bad Request.
		fmt.Fprintf(w, err.Error())
		return
	}

	// validate the JSON data fields
	validate := validator.New()
	if err = validate.Struct(smsContext); err != nil {
		log.Printf("All the required fields are not available, %v", err)
		w.WriteHeader(400) // Return 400 Bad Request.
		fmt.Fprintf(w, err.Error())
		return
	}
	log.Println(smsContext)

	//send SMS function
	_, err = myCustomSmsFunction(smsContext)
	if err != nil {
		log.Printf("Issue with the SMS module, %v", err)
		w.WriteHeader(500) // Return 500 Internal Server Error.
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(200)
	fmt.Fprintf(w, "SMS Sent successfully.")
}

var n int = 0

func myCustomSmsFunction(s smsSchema) (bool, error) {
	n++
	// mock, if the SIM is not available, send error
	if n%2 == 0 {
		return false, errors.New(fmt.Sprint("SIM Card not available"))
	}

	//on successful sms sent
	return true, nil
}

func main() {
	http.HandleFunc("/sms", sms)

	http.ListenAndServe(":8090", nil)
}
