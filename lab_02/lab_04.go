package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a buffered channel with size 2
	messages := make(chan string, 2)

	messages <- "First message"
	messages <- "Second message"

	go func() {
		select {
		case messages <- "Third message":
			fmt.Println("Third message sent!")
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout: Could not send third message - channel buffer is full")
		}
	}()

	// Give the goroutine time to execute
	time.Sleep(3 * time.Second)

	// Read from the channel
	fmt.Println("\nReading messages from channel:")
	msg1 := <-messages
	fmt.Println("Received:", msg1)

	// After reading one message, the goroutine should be able to send its message
	time.Sleep(1 * time.Second)

	msg2 := <-messages
	fmt.Println("Received:", msg2)

	// Try to read the third message if it was sent
	select {
	case msg3 := <-messages:
		fmt.Println("Received:", msg3)
	default:
		fmt.Println("No more messages in channel")
	}
}
