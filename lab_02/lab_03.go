package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string) // Create a channel

	go func() {
		time.Sleep(2 * time.Second)
		messages <- "hello from 1" // Send a message to the channel
	}()

	go func() {
		time.Sleep(1 * time.Second)
		messages <- "hello from 2" // Send a message to the channel
	}()

	go func() {
		time.Sleep(3 * time.Second)
		messages <- "hello from 3" // Send a message to the channel
	}()

	fmt.Println("Main started")

	msg1 := <-messages
	fmt.Println(msg1)

	msg2 := <-messages
	fmt.Println(msg2)

	msg3 := <-messages
	fmt.Println(msg3)

	fmt.Println("Main finished")
}
