package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	// Get a greeting message and print it.

	message, err := greetings.Hello(text)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
