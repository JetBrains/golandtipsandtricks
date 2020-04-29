package main

import (
	"fmt"
)

func main() {
	// Step 1. Method-like completion helps you quickly find
	//  functions that match their first parameter to the
	//  type you invoke it for
	// Shortcut: Ctrl + Space + Space on Windows/Linux
	//           ^ + Space + Space on macOS

	msg := "hello GoPHeRS!"

	// E.g. you can use ToTitle quickly from here
	msg = msg.
		fmt.Println(msg)
}
