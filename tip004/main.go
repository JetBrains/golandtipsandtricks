package main

import (
	"fmt"
)

func main() {
	// Postfix completion helps you quickly wrap code
	// with predefined code templates.

	// E.g. Use .var after the below expression
	sendMessage("hello GoPHeRS!")
}

func sendMessage(msg string) (int, error) {
	return fmt.Println(msg)
}
