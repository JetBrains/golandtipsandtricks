package main

func main() {
	// Parameter Info helps you to see parameters in a function/method.
	// Step 1. Invoke it here
	myFunc()

	// But also in a struct.
	// Step 2. Invoke it here
	_ = demo{}
}

type (
	SubStruct struct {
		Field1 int
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
