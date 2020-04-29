package main

import (
	"encoding/json"
	"fmt"
)

// Step 0. First go to 1. then 2.

// Step 2. Use multi-cursor mode to select both fields and then
//  type ` and use "json" Live Template
// Shortcut: Alt + Shift + Insert on Windows/Linux
//           ⌘ + Shift + 8 on macOS

func main() {

	// Step 1. Invoke creating a type here (Alt+Enter)
	//  then adding the missing fields to it
	msg := message{
		Field1: 1,
		Field2: "2",
	}

	out, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
}
