package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Set seed for random number generation
	rand.Seed(time.Now().UnixNano())

	// Create two channels
	channel1 := make(chan string)
	channel2 := make(chan string)

	fmt.Println("Starting goroutines...")

	// Launch first goroutine
	go func() {
		// Random duration between 1-5 seconds
		duration := time.Duration(1+rand.Intn(5)) * time.Second
		time.Sleep(duration)
		channel1 <- fmt.Sprintf("Message from channel 1 (after %v)", duration)
	}()

	// Launch second goroutine
	go func() {
		// Random duration between 1-5 seconds
		duration := time.Duration(1+rand.Intn(5)) * time.Second
		time.Sleep(duration)
		channel2 <- fmt.Sprintf("Message from channel 2 (after %v)", duration)
	}()

	// Use select to wait for whichever message comes first
	select {
	case message1 := <-channel1:
		fmt.Println(message1)
	case message2 := <-channel2:
		fmt.Println(message2)
	}

	fmt.Println("First message received, program complete!")
}
