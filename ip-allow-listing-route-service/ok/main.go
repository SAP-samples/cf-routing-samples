package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := http.ListenAndServe(":"+os.Getenv("PORT"), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, "ok")
		if err != nil {
			fmt.Printf("error: handle request: %s\n", err.Error())
		}
	}))
	if err != nil {
		fmt.Printf("error: server exited: %s\n", err.Error())
	}
}
