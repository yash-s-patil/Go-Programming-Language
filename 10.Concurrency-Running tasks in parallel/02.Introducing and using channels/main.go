package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func main() {
	// go greet("Nice to meet you!")
	// go greet("How are you?")

	// channel is basically a value, that is transmitted as data, that can be used as a communication channel when working with goroutines
	done := make(chan bool)

	go slowGreet("How ... are ... you ...?", done)
	// make the data come out of channel
	<-done
	// go greet("I hope you're liking the course!")
}
