package main

type demoType struct {
	Field1 int
}

// Step 0. First go to step 1, then show how auto-testing works.

// Step 1. Invoke Generate action here and then use the Test Function feature
// Shortcut: Alt + Insert on Windows/Linux
//           ⌘ + N on macOS

func (d *demoType) AddToField(val int) int {
	d.Field1 += val
	return d.Field1
}

func main() {
	_ = demoType{}
}
