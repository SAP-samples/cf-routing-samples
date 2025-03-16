package main

import (
	"fmt"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

var requests atomic.Int32

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /delay/{delay}", delayHandler)
	mux.HandleFunc("GET /requests", requestsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+os.Getenv("PORT"), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests.Add(1)
		defer requests.Add(-1)

		mux.ServeHTTP(w, r)
	}))
}

func delayHandler(w http.ResponseWriter, r *http.Request) {
	d, err := time.ParseDuration(r.PathValue("delay"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error: failed to parse duration: %s", err.Error())
		return
	}

	time.Sleep(d)

	w.WriteHeader(http.StatusOK)
}

func requestsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%d", requests.Load())
}
