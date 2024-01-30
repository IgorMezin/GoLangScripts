package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	// A slice of message formats.
	formats_start := []string{
		"Hi", "Heil", "Hello", "Hallo",
	}
	formats_end := []string{
		"sleepy", "basic", "milky", "full of",
	}
	return formats_start[rand.Intn(len(formats_start))] + ", " + formats_end[rand.Intn(len(formats_end))] + "%v"
}
