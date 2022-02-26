package main

import (
	"log"
	"net/http"
	"os"
)

var (
	secret = os.Getenv("SECRET_PASS_PHRASE")
)

func main() {
	if secret == "" {
		log.Fatal("SECRET_PASS_PHRASE is blank")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Speak friend and enter."))
	})

	http.HandleFunc("/"+secret, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the secret clubhouse."))
	})

	http.ListenAndServe(":8080", nil)
}
