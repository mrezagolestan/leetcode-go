package main

import (
	"leetcode-go/dependency"
	"log"
	"net/http"
)

func MyGreeterResponseHandler(w http.ResponseWriter, R *http.Request) {
	dependency.Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterResponseHandler)))
}
