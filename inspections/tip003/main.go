package main

import (
	"context"
	"net/http"
)

type demoType string

const isDemo demoType = "yes"

// Step 1. Invoke the Navigate to field with the same json tag action
// Shortcut: Alt + Enter on Windows/Linux
//           ⌥ + Enter on macOS

func demoHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if (ctx.Value(isDemo)).(demoType) == isDemo {
		cancel()
	}

	// Step 2. Uncomment this to see the inspection in action
	// cancel()
}

func main() {}
