package main

import "fmt"

func main() {
	// Completion doesn't work just in source code.
	// For example, you can set a breakpoint condition and use completion

	for i := 0; i < 42; i++ {
		msg, err := doStuff(i)
		// E.g. Set a breakpoint below and use completion for err != nil
		if err != nil {
			break
		}
		fmt.Println(msg)
	}
}

func doStuff(n int) (string, error) {
	if n == 23 {
		return "", fmt.Errorf("ohhhh.... my")
	}
	return fmt.Sprintf("%d", n), nil
}
