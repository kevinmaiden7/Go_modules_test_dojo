package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Alice", "Bob"}

	// Request a greeting message.
	// message, err := greetings.Hello("Kevin")

	// Request a greeting message.
	message, err := greetings.Hellos(names)

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err) // Print the error and stop the program
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)

	fmt.Println(greetings.HelloPhrase())
}
