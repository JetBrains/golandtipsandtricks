package main

import "io"

func main() {
	// Step 1. Invoke Type Hierarchy on the Reader type.
	// Shortcut: Ctrl + H on Windows/Linux
	//           ^ + H on macOS

	// Step 2. Invoke Call Hierarchy on the Read method.
	// Shortcut: Ctrl + Alt + H on Windows/Linux
	//           ⌥ + ^ + H on macOS

	// Step 3. Use Quick Definition on any types to preview them
	// Shortcut: Ctrl + Shift + Space on Windows/Linux
	//           ⌥ + Space on macOS

	_, _ = io.Reader(nil).Read(nil)
}
