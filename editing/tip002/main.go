package main

func main() {
	// Parameter Info helps you to see parameters in a function/method.
	// Shortcut: Ctrl + P on Windows/Linux
	//           ⌘ + P on macOS

	// Step 1. Invoke it here
	myFunc()

	// But also in a struct.
	// Step 2. Invoke it here
	_ = demo{}

	// Step 3. Or here
	_ = demo{0, "", SubStruct{}}
}

type (
	SubStruct struct {
		SubField1 int
	}

	demo struct {
		Field1 int
		Field2 string
		SubStruct
	}
)

func myFunc(n int, s string) (int, string) {
	_, _ = n, s

	return 1, "demo"
}
