package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

var durations = []time.Duration{
	time.Millisecond,
	time.Millisecond,
	10 * time.Second,
}

func main() {
	c := http.DefaultClient

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s TARGET_URL\n", os.Args[0])
		fmt.Println("error: please provide the target URL as first argument")
		os.Exit(1)
	}

	target, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Printf("error: parse target URL: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("starting 30 requests")

	wg := &sync.WaitGroup{}
	for i := range 30 {
		u := *target
		u.Path = fmt.Sprintf("/delay/%s", durations[rand.Intn(len(durations))].String())

		wg.Add(1)
		go func() {
			defer wg.Done()
			res, err := c.Do(&http.Request{
				Method: http.MethodGet,
				URL:    &u,
			})
			if err != nil {
				fmt.Printf("request failed: %s\n", err.Error())
				return
			}
			if res.StatusCode != http.StatusOK {
				fmt.Printf("request failed: unsuccessful status code %s\n", res.Status)
				return
			}
		}()

		fmt.Printf("\r%d / 30", i+1)

		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("\nwaiting for all requests to finish")

	wg.Wait()
}
