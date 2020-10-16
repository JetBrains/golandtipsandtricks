package main

// Step 0. Go to main_test.go

type demoType struct {
	Field1 int
}

func (d *demoType) AddToField(val int) int {
	d.Field1 += val
	return d.Field1
}

func main() {
	_ = demoType{}
}
