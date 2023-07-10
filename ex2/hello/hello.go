package main

import (
	"fmt"
	"log"

	"ex2/greetings"
)

/*
 * In Go, code executed as an application must be in a "main" package
 */
func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("Gladys")
	// If an error was returned, print it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returend, print the returend message
	// to the console.
	fmt.Println(message)
}
