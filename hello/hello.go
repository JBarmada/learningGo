package main

import (
	"log"
	"fmt"

	"github.com/JBarmada/LearningGo/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
	log.SetPrefix("logger of geetings: ")
	log.SetFlags(0)

	// call hello function (request a greeting)
	message, err := greetings.Hello("Jad")
	// check for an error 
	if err != nil {
		log.Fatal(err)
	}
	// get a greeting and print it <- will not reach due to Fatal function
	fmt.Println(message)
}