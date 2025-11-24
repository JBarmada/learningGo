package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

// hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// check that a name was given
	if name == "" {
		// no name given, return an error
		return name, errors.New("empty name") 
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randGreeting(), name)
	return message, nil
}

func randGreeting() string {
	// a slice of greetings
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
