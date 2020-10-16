package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	_ = r
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte("Hello World from the Home page"))
	time.Sleep(time.Duration(50+rand.Intn(100)) * time.Millisecond)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	_ = r
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte("Login page"))
	time.Sleep(time.Duration(50+rand.Intn(100)) * time.Millisecond)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	_ = r
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte("Login page"))
	time.Sleep(time.Duration(50+rand.Intn(100)) * time.Millisecond)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	_ = r
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte("Products page"))
	time.Sleep(time.Duration(50+rand.Intn(100)) * time.Millisecond)
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	rVars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	_, _ = fmt.Fprintf(w, "Product %s page", rVars["productID"])
	time.Sleep(time.Duration(50+rand.Intn(100)) * time.Millisecond)
}

func basketHandler(w http.ResponseWriter, r *http.Request) {
	_ = r
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte("Basket page"))
	time.Sleep(time.Duration(50+rand.Intn(100)) * time.Millisecond)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	_ = r
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte("About page"))
	time.Sleep(time.Duration(50+rand.Intn(100)) * time.Millisecond)
}
