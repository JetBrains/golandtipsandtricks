package main

// Step 1. Invoke the Extract Interface refactoring on demo
//  and extract To File main (without .go extension)

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
