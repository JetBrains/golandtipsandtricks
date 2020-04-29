package main

import (
	"fmt"
	"strconv"
)

func convertToString(n int) string {
	return fmt.Sprintf("%d", n)
}

func main() {
	// Step 1. Cyclic Expand Word completion helps quickly using identifiers
	//  from above, or below the current caret position
	// Shortcut: Alt + / on Windows/Linux
	//           ⌥ + / on macOS

	// Step 2. Cyclic Expand Word (Backward) (Alt+Shift+/) also works
	// Shortcut: Alt + Shift + / on Windows/Linux
	//           ⌥ + Shift + / on macOS

	// E.g. At the end of "con" use Cyclic Expand Word
	fmt.Printf("running con")
}

func convertToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func conditionalOutput(s string) {
	if s == "yes" {
		fmt.Println("Yes!")
	} else {
		fmt.Println("No!")
	}
}
