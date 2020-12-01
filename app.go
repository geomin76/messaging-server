package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// type EmailBody struct {
// 	toEmail string `json:"toEmail"`
// 	msg     string `json:"msg"`
// 	from    string `json:"from"`
// }

func email(w http.ResponseWriter, r *http.Request) {
	// Getting request body
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Printf("Error reading body: %v", readErr)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	fmt.Println(string(body))

	// Sending email
	auth := smtp.PlainAuth("", os.Getenv("EMAIL"), os.Getenv("EMAIL_PASS"), "smtp.gmail.com")
	to := []string{"geomin76@gmail.com"}
	msg := []byte("Hello, World!")

	err := smtp.SendMail("smtp.gmail.com:587", auth, os.Getenv("EMAIL"), to, msg)
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
