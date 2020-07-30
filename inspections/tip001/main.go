package main

import (
	"fmt"
)

func main() {
	d := uint64(11111111111111111110)
	fmt.Printf("Initial code:\t%s\n", string(d))

	// Step 1. Invoke the Convert integer to string quickfix
	// Shortcut: Alt + Enter on Windows/Linux
	//           ⌥ + Enter on macOS

	fmt.Printf("Correct code:\t%s\n", string(d))
}
