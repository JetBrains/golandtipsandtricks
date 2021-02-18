package main

import (
	"fmt"
)

func main() {
	// Code completion works in language injections too!

	// Step 1. Inject (Alt+Enter | Inject Language | JSON) below
	json := `{
  "field": "some value"
}`
	fmt.Println(json)
}
