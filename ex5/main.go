package main

import (
	"fmt"
)

/*
 * A function, that accepts a channel as a parameter
 * - Prints to stdout
 * - Sends a message to the channel
 */
func printMessage(ch chan string) {
	fmt.Println("Hello, from our printMessage")
	ch <- "Sending to channel...from printMessage"
}

func main() {
	fmt.Println("EX5 - Context")

	// Let's make a channel
	ch := make(chan string)

	// And we're making an inline goroutine
	go func() {
		fmt.Println("Hello, from our inline goroutine")
		// ...let's send a message to our channel
		ch <- "Sending to channel...from inline goroutine"
	}()

	go printMessage(ch)
	fmt.Println("Hello, from our main()")

	// Let's read from our channel
	message := <-ch
	fmt.Println("from the channel: ", message)

	// Throwaway the second value because...happy birthday to the ground
	<-ch
}
