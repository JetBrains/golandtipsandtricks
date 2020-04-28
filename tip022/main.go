package main

import "io"

func _(v io.Reader)  {
	// Step 1. Check if v implements the io.Reader interface
	//  by using completion on v.(|)
	_, _ = v.()
}

type myType struct{}

func (myType) Read(p []byte) (n int, err error) {
	panic("implement me")
}

func (myType) Write(p []byte) (n int, err error) {
	panic("implement me")
}

func (myType) Close() error {
	panic("implement me")
}
