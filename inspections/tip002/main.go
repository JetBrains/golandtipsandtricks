package main

// Step 1. Invoke the Navigate to field with the same json tag action
// Shortcut: Alt + Enter on Windows/Linux
//           ⌥ + Enter on macOS

type user struct {
	Username string `json:"name"`
	FirstName string `json:"name"`
	LastName string `json:"name"`
}

func main() {
	_ = user{}
}
