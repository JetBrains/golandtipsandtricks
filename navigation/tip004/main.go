package main

import (
	"fmt"
	"time"
)

// longCall uses fmt.Printf to print the data supplied to it
// You can always call longCall to print this data.
// Don't use this for debugging however, GoLand has better debugging
// functionality for this if needed.
func longCall(name, surname, email string, dateOfBirth time.Time, username, password string) {
	fmt.Printf(
		"%s %s (%s) - born on %s with the username: %s and password %s\n",
		name,
		surname,
		email,
		dateOfBirth,
		username,
		password,
	)
}

func main() {
	name := "Patan"
	surname := "Florin"
	email := "florin [] jetbrains.com"
	dateOfBirth := time.Date(
		2020,
		1,
		1,
		0,
		0,
		0,
		0,
		time.UTC,
	)
	username := "dlsniper"
	password := "GoLand!"

	longCall(name, surname, email, dateOfBirth, username, password)
}
