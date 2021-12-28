package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type payload struct {
	Headers map[string][]string `json:"headers"`
	Proto   string              `json:"protocol"`
}

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "Hello! This Go application is speaking plain text HTTP2 (H2C) with the CF routing layer")
	})

	h2s := &http2.Server{}

	port := os.Getenv("PORT")
	if port == "" {
		panic("environment variable PORT not found")
	}

	h1s := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: h2c.NewHandler(handler, h2s),
	}

	log.Fatal(h1s.ListenAndServe())
}
