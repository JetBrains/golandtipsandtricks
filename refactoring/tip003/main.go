package main

// Step 1. Invoke the Extract Interface refactoring
//  on demo and extract an interface named demoer
//  to file main (without .go extension)
// Shortcut: Ctrl + Alt + Shift + T on Windows/Linux
//           ^ + T on macOS

type demo struct{}

func (demo) Method1() int {
	return 0
}

func (demo) Method2() (string, error) {
	return "", nil
}

func main() {
	_ = demo{}
	_ = demoer(nil)
}
