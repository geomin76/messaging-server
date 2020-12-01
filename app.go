package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func email(w http.ResponseWriter, r *http.Request) {
	auth := smtp.PlainAuth("", "email", "pass", "smtp.gmail.com")

	// Here we do it all: connect to our server, set up a message and send it
	to := []string{""}
	msg := []byte("Hello, World!")

	err := smtp.SendMail("smtp.gmail.com:587", auth, "from", to, msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Email sent")
}

func text(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Text sent")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/email", email)
	http.HandleFunc("/text", text)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("App is running at http://localhost:10000")
	handleRequests()
}
