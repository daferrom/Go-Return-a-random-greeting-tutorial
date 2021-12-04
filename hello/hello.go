// Declare the main package (A group of functions )
package main

//fmt package contains functions for formatting text
import (
	"fmt"
	"log"

	"example.com/greetings"
)

// main function executes by default when you run the main package./

// main function is to print a message on console
//A main function executes by default when you run the main package./
func main() {
	//Set the properties of the predifined logger, including
	//the log entry prefix and a flag to disable printing
	//the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//Request a greeting message.
	message, err := greetings.Hello("Diego")
	// If an error was returned, print it to the console and
	// exit the program
	if err != nil {
		log.Fatal(err)
	}
	//If no error was returned, print the returned message
	// to the console
	fmt.Println(message)
}
