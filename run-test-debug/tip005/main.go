package main

import "fmt"

func main() {
	// Step 1. Invoke the debugger and then use Smart Step Into
	//  to choose between makeDemo() and Greet() to step into
	// Shortcut: Shift + F7 on Windows/Linux
	//           Shift + F7 on macOS

	message := "各位 Gopher 晚上好"
	makeDemo().Greet(message)
}

type demo struct{
	init string
}

func (d demo) DebugString() string {
	// Step 2. Mention that DebugString() performs magic in the debugger
	return d.init
}

func (d demo) Greet(msg string) {
	fmt.Println(msg)
}

func makeDemo() demo {
	return demo{"MyTypeIsInitialized"}
}
