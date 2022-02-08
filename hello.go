package main

import (
	"fmt"

	"github.com/alexmacartney/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Alex")
	fmt.Println(message)
}
