package main

import (
	"fmt"
	"time"
)

func main() {
	// Step 1. Type the date and time in a format like this:
	//         YYYY-MM-DD hh:mm:ss zone
	now := time.Now().Format("")
	fmt.Println(now)

	// Step 2. It works here too
	_, _ = time.Parse("", now)
}
