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

func MultiHello(names []string) (map[string]string, error){
	// a map to associate names with messages
	message := make(map[string]string)

	// go thru the slice of names, and get a message for each name
	for _, name :=  range names {
		greeting, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// map the greeting to the name
		message[name] = greeting
	}
	return message, nil
}

// randGreeting returns a greeting chosen at random from a set of greeting messages.
func randGreeting() string {
	// a slice of greetings
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
