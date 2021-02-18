package main

import (
	"fmt"
)

func main() {
	// Postfix completion helps you quickly wrap code
	// with predefined code templates.

	// E.g. Use .var after the below expression
	_, err := sendMessage("hello GoPHeRS!")
	if err != nil {
		panic(err)
	}
}

func sendMessage(msg string) (int, error) {
	return fmt.Println(msg)
}
