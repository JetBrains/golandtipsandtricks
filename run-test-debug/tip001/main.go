package main

import "fmt"

func main() {
	// Completion doesn't work just in source code.
	// For example, you can set a breakpoint condition
	// and use completion to quickly create the breakpoint

	for i := 0; i < 42; i++ {
		msg, err := doStuff(i)

		// Step 1. Set a breakpoint below
		// Shortcut: Ctrl + F8 on Windows/Linux
		//           ⌘ + F8 on macOS

		// Step 2. Use completion for err != nil in breakpoint condition
		// Shortcut: Ctrl + Shift + F8 on Windows/Linux
		//           ⌘ + Shift + F8 on macOS
		if err != nil {
			break
		}
		fmt.Println(msg)
	}
}

func doStuff(n int) (string, error) {
	if n == 23 {
		return "", fmt.Errorf("ohhhh.... my")
	}
	return fmt.Sprintf("%d", n), nil
}
