package main

import (
	"fmt"
	"strconv"
)

type subDemoStruct struct {
	SubField1 string
	SubField2 int
}

type demoStruct struct {
	Field1 int
	Field2 string
	Field3 subDemoStruct
}

func main() {
	mySum := "Let's sum a string and an int!"
	a := "string"
	b := 1
	// Step 1. Set a breakpoint on the line below
	// Shortcut: Ctrl + F8 on Windows/Linux
	//			 Cmd + F8 on macOS
	// Step 2. Run the debugger
	// Step 3. You can right-click on a symbol, such as mySum,
	// 		   and use the "Add Inline Watch" action
	// Step 4. Step into the next line to observe it.
	mySum = sumThese(a, b)
	fmt.Printf("mySum is = %s\n", mySum)

	random := demoStruct{
		Field1: 1,
		Field2: "string",
		Field3: subDemoStruct{
			SubField1: "subString",
			SubField2: 2,
		},
	}
	// Step 5. Observe how to expand the random symbol
	// 		   using the inline debugger values.
	fmt.Printf("random is = %v\n", random)
}

func sumThese(s string, i int) string {
	return s + strconv.Itoa(i)
}
