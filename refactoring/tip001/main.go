package main

// Implementing an interface is usually complex and repetitive.
// So let GoLand do it for you.

// Step 1. Invoke "Generate" here and use
//  Implement Interface to create a new type then implement
//  any interface, e.g. rewrcl ;)
// Shortcut: Alt + Insert on Windows/Linux
//           ⌘ + N on macOS

type demoType struct{}

func (demoType) Read(p []byte) (n int, err error) {
	panic("implement me")
}

func (demoType) Write(p []byte) (n int, err error) {
	panic("implement me")
}

func (demoType) Close() error {
	panic("implement me")
}

func main() {

}
