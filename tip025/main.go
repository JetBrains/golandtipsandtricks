package main

func addNumbers(a, b int) int {
	return a + b
}

func main() {
	// Step 1. Invoke the change signature refactoring
	// Step 2. Move function to a different package

	_ = addNumbers(1, 2)
}
