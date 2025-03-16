package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	err := Main()
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}
}

func Main() error {
	c := http.DefaultClient

	if len(os.Args) < 2 {
		return fmt.Errorf("please provide the target URL as first argument")
	}

	target, err := url.Parse(os.Args[1])
	if err != nil {
		return err
	}

	target.Path = "/requests"

	id, err := appGuid()
	if err != nil {
		return err
	}

	fmt.Printf("using app GUID '%s'\n", id)
	// For cursor alignment.
	fmt.Printf("%s\nInstance 1: 1\nInstance 2: 1\nInstance 3: 1", time.Now().Format("2006-01-02 15:04:05"))

	var instances [3]int
	for {
		for i := range instances {
			res, err := c.Do(&http.Request{
				Method: http.MethodGet,
				URL:    target,
				Header: http.Header{
					"X-Cf-App-Instance": []string{id + ":" + strconv.Itoa(i)},
				},
			})
			if err != nil {
				return err
			}

			b, err := io.ReadAll(res.Body)
			if err != nil {
				return err
			}

			instances[i], err = strconv.Atoi(string(b))
			if err != nil {
				return err
			}
		}

		fmt.Printf("\r\x1b[3A%s\nInstance 1: %d\nInstance 2: %d\nInstance 3: %d", time.Now().Format("2006-01-02 15:04:05"), instances[0], instances[1], instances[2])

		time.Sleep(time.Second)
	}
}

func appGuid() (string, error) {
	cmd := exec.Command("cf", "app", "--guid", "per-route-loadbalancing-demo")

	b, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("run 'cf app': %w\n%s", err, string(b))
	}

	return strings.TrimSpace(string(b)), nil
}
