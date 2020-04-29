package main

import "fmt"

func main() {
	// Step 1. Invoke the debugger and then use Smart Step Into
	//  to choose between makeDemo() and Greet() to step into

	message := "各位 Gopher 晚上好"
	makeDemo().Greet(message)
}

type demo struct{}

func (d demo) Greet(msg string) {
	fmt.Println(msg)
}

func makeDemo() demo {
	return demo{}
}
