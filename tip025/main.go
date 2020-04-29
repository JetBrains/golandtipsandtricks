package main

func main() {
	// Step 1. Invoke the Change Signature refactoring
	// Shortcut: Ctrl + Alt + Shift + T on Windows/Linux
	//           ^ + T on macOS

	// Step 2. Move function to a different package
	// Shortcut: Ctrl + Alt + Shift + T on Windows/Linux
	//           ^ + T on macOS

	_ = addNumbers(1, 2)
}

func addNumbers(a, b int) int {
	return a + b
}
