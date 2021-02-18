package main

import "net/http"

func main() {
	// Smart Type Completion helps narrowing down the values in a scope
	// Shortcut: Ctrl + Shift + Space on Windows/Linux
	//           ^ + Shift + Space on macOS

	// Step 1. In the http.HandleFunc() call below, Smart Type Completion
	//  can complete the expected lambda function correctly.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
}
