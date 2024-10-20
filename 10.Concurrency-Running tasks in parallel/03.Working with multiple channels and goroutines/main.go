package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan) // now we cant send data anymore
	// once the operation is done we dont need the channel anymore
}

func main() {
	// channel is basically a value, that is transmitted as data, that can be used as a communication channel when working with goroutines
	// channel is used to wait for the completion of goroutine
	done := make(chan bool)

	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)
	// make all the data come out of channel
	for range done {
		// fmt.Println(doneChan)
	}

}

// race condition - We dont wait for all goroutines to finish, the function that completes first, simply ends the execution of the program, as soon as we got one value out of the channel
// for this we have to make all the data for all goroutines function come out of the channel
