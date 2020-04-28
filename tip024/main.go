package main

import "io"

func main() {
	// Step 1. Invoke Type Hierarchy on the Reader type.
	// Step 2. Invoke Call Hierarchy on the Read method.

	_, _ = io.Reader(nil).Read(nil)
}
