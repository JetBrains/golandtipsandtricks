// For more details, read https://blog.jetbrains.com/go/2020/03/03/how-to-find-goroutines-during-debugging/

package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/dlsniper/debugger"
	"github.com/gorilla/mux"
)

// Step 1. Run the code below using the debug mode

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("initializing the server...")

	m := mux.NewRouter()
	m.HandleFunc("/", debugger.Middleware(homeHandler, func(r *http.Request) []string {
		return []string{
			"request", "real",
			"path", r.RequestURI,
		}
	}))
	m.HandleFunc("/login", debugger.Middleware(loginHandler, func(r *http.Request) []string {
		return []string{
			"request", "real",
			"path", r.RequestURI,
		}
	}))
	m.HandleFunc("/logout", debugger.Middleware(logoutHandler, func(r *http.Request) []string {
		return []string{
			"request", "real",
			"path", r.RequestURI,
		}
	}))
	m.HandleFunc("/products", debugger.Middleware(productsHandler, func(r *http.Request) []string {
		return []string{
			"request", "real",
			"path", r.RequestURI,
		}
	}))
	m.HandleFunc("/product/{productID}", debugger.Middleware(productHandler, func(r *http.Request) []string {
		return []string{
			"request", "real",
			"path", r.RequestURI,
		}
	}))
	m.HandleFunc("/basket", debugger.Middleware(basketHandler, func(r *http.Request) []string {
		return []string{
			"request", "real",
			"path", r.RequestURI,
		}
	}))
	m.HandleFunc("/about", debugger.Middleware(aboutHandler, func(r *http.Request) []string {
		return []string{
			"request", "real",
			"path", r.RequestURI,
		}
	}))

	srv := http.Server{
		Addr:              "127.0.0.1:8080",
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       15 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    200 * 1000,
	}

	go fakeTraffic()

	// Late binding of our shutdown path so that we don't call it from the traffic faker goroutine
	m.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		_ = srv.Shutdown(context.Background())
	})

	srv.Handler = m

	log.Println("starting the server...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
