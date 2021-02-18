package main

type demoType struct {
	Field1 int
}

// Step 1. Completion works in comments as well

// E.g. Add a comment via quickfix and then use completion for Field1

//AddToField something like this demoType
func (d *demoType) AddToField(val int) int {
	d.Field1 += val
	return d.Field1
}

func main() {
	_ = demoType{}
}
