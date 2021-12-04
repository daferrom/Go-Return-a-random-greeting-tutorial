//Declare a greetings package to collect related functions.
package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//Hello returns a greeting for a named person.
//This function takes a name parameter whose type is string.
//The function also returns a string.
func Hello(name string) (string, error) {
	//If no name was given , return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	//Create a message using a rnadom format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

//randomFormat returns one of a set of a greeting messages.
//The returned message is selected at random.
func randomFormat() string {
	// A slice of formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	//Return a randomly selected message format specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}