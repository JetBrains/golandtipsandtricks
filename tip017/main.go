package main

type demoType struct {
	Field1 int
}

func (d *demoType) AddToField(val int) int {
	// Step 1. Use Add Selection for Next Occurrence for Field1

	d.Field1 += val
	return d.Field1
}

func main() {
	_ = demoType{
		Field1: 1,
	}
}
