package main

import (
	"leetcode-go/dependency"
	"leetcode-go/mocking"
	"net/http"
	"os"
	"time"
)

func MyGreeterResponseHandler(w http.ResponseWriter, R *http.Request) {
	dependency.Greet(w, "world")
}

func main() {
	//log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterResponseHandler)))

	sleeper := &mocking.ConfigurableSleeper{1 * time.Second, time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
