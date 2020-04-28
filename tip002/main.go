package main

import "net/http"

func main() {
	// Smart Type Completion helps narrowing down the values in a scope

	// E.g. In the http.HandleFunc() call below, Smart Type Completion
	//  can complete the expected lambda function correctly.
	http.HandleFunc("/")
}
