package main

import (
	"time"
)

// Step 1. Invoke Expand signature types then
//         invoke Reuse signature types in the function definition below
// Shortcut: Alt + Enter on Windows/Linux
//           ⌥ + Enter on macOS

func longCall(name, surname, email string, dateOfBirth time.Time, username, password string) {
	_, _, _, _, _, _ = name, surname, email, dateOfBirth, username, password
}

func main() {
	name := "Patan"
	surname := "Florin"
	email := "florin [] jetbrains.com"

	// Step 2. Invoke Put arguments on a separate line then
	//         invoke Put arguments on one line
	// Shortcut: Alt + Enter on Windows/Linux
	//           ⌥ + Enter on macOS

	dateOfBirth := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	username := "dlsniper"
	password := "GoLand!"

	longCall(name, surname, email, dateOfBirth, username, password)

	// Step 3. Invoke Settings/Preferences and search for Wrapping and Braces
	//         Change the preferences under Go Code Style
	// Shortcut: Ctrl + Alt + S on Windows/Linux
	//           ⌘ + , on macOS
}
