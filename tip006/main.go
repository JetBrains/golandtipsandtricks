package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	// Step 1. Live templates are useful to quickly predefine code or
	//  complete repeatable code.

	// E.g. when you want to connect to a database, you can use "dbds"
	dsn :=
	_, err := sql.Open("postgres", dsn)
	
	// Step 2. Use the if val != nil template to handle the error

}

func _() (int, error) {
	// Step 3. Use the return template for returning an in and a nil

}
