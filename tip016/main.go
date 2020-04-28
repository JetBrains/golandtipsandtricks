package main

type demoType struct {
	Field1 int
}

func (d *demoType) AddToField(val int) int {
	// Step 1. Use Expand Selection
	// Step 2. Use Shrink Selection

	d.Field1 += val
	return d.Field1
}

func main() {
	_ = demoType{}
}
