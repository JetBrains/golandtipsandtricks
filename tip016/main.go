package main

type demoType struct {
	Field1 int
}

func (d *demoType) AddToField(val int) int {
	// Step 1. Use Extend Selection (Ctrl+W)
	// Shortcut: Ctrl + W on Windows/Linux
	//           ⌥ + Up arrow on macOS
	// Step 2. Use Shrink Selection (Ctrl+Shift+W)
	// Shortcut: Ctrl + Shift + W on Windows/Linux
	//           ⌥ + Down arrow on macOS

	d.Field1 += val
	return d.Field1
}

func main() {
	_ = demoType{}
}
