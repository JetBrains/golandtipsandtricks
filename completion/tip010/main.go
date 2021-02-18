package main

import (
	"database/sql"
	"fmt"
	"time"
)

func main() {
	// Step 1. Type the date and time in a format like this:
	//         YYYY-MM-DD hh:mm:ss zone
	now := time.Now().Format("2006-01-002")
	fmt.Println(now)

	// Step 2. It works here too
	_, _ = time.Parse("06-Jan-2", now)

	db, _ := sql.Open("", "")
	const query = "SELECT jetbrains.ides.name FROM jetbrains.ides"
	db.Query(query)
}
