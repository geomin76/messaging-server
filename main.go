package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/rs/cors"
	"github.com/sfreiberg/gotwilio"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// setupCorsResponse(&w, r)
	fmt.Fprintf(w, "Hello, World!")
}

type Email struct {
	ToEmail string
	Msg     string
	Subject string
	From    string
}

type Text struct {
	From string
	Msg  string
	To   string
}

// func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// 	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
// }

func email(w http.ResponseWriter, r *http.Request) {
	// setupCorsResponse(&w, r)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	// Getting request body
	var email Email

	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Printf("Error reading body: %v", readErr)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	// Parsing request.body to Email Struct
	json.Unmarshal([]byte(string(body)), &email)

	// Sending email
	auth := smtp.PlainAuth("", os.Getenv("EMAIL"), os.Getenv("EMAIL_PASS"), "smtp.gmail.com")
	to := []string{string(email.ToEmail)}
	msg := []byte("To: " + string(email.ToEmail) + "\r\n" +
		"Subject: " + string(email.Subject) + "\r\n" +
		"\r\n" +
		"From: " + string(email.From) + "\r\n\n" +
		"Message: " + string(email.Msg) + "\r\n")

	err := smtp.SendMail("smtp.gmail.com:587", auth, os.Getenv("EMAIL"), to, msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Email sent")
}

func text(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	var text Text

	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Printf("Error reading body: %v", readErr)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	// Parsing request.body to Email Struct
	json.Unmarshal([]byte(string(body)), &text)

	accountSid := string(os.Getenv("PHONE_KEY"))
	authToken := string(os.Getenv("PHONE_SECRET"))
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := string(os.Getenv("TWILIO_NUMBER"))
	to := string(text.To)
	message := "From: " + string(text.From) + "\r\n\n" +
		"Message: " + string(text.Msg)
	twilio.SendSMS(from, to, message, "", "")

	fmt.Fprintf(w, "Text sent with Twilio")
}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/email", email)
	mux.HandleFunc("/text", text)

	handler := cors.Default().Handler(mux)
	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		panic(err)
	}
}
