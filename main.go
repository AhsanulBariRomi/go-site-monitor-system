package main

import (
	"fmt"
	"net/http"
	"time"
)

// Result struct to hold the status of a check
type Result struct {
	URL     string
	Status  string
	Latency time.Duration
}

// checkUrl visits a URL and sends the result to a channel
func checkUrl(url string, c chan Result) {
	start := time.Now()
	resp, err := http.Get(url)
	latency := time.Since(start)

	if err != nil {
		c <- Result{URL: url, Status: "DOWN", Latency: latency}
		return
	}
	defer resp.Body.Close()

	c <- Result{URL: url, Status: fmt.Sprintf("%d %s", resp.StatusCode, http.StatusText(resp.StatusCode)), Latency: latency}
}

func main() {
	// The list of websites we want to test
	urls := []string{
		"https://www.google.com",
		"https://www.clyso.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://www.tu-dresden.de",
		"https://shouldnotwork.com",
	}

	// Create a channel to communicate between routines
	c := make(chan Result)

	fmt.Println("Starting Concurrent Health Check...")

	// Spin up a Goroutine for each URL
	for _, url := range urls {
		go checkUrl(url, c)
	}

	// Collect results from the channel
	for i := 0; i < len(urls); i++ {
		result := <-c
		if result.Status == "DOWN" {
			fmt.Printf("[DEAD] %s is DOWN (Error)\n", result.URL)
		} else {
			fmt.Printf("[ALLIVE] %s is UP | Status: %s | Latency: %v\n", result.URL, result.Status, result.Latency)
		}
	}

	fmt.Println("All checks completed.")
}
