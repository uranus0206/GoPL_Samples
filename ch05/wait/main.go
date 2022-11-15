package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "https://googl.com"

	if err := WaitForServer(url); err != nil {
		log.Fatalf("Site is down: %v\n", err)
	}
}

// WaitForServer attempts to contact the server of a URL.
// IT tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
func WaitForServer(url string) error {
	log.SetPrefix("wait: ")
	log.SetFlags(0)
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)

	// Retry before deadline
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success return
		}
		sleepTime := time.Second << uint(tries)
		log.Printf("server not responding (%s); sleep: %v, retrying...", err, sleepTime)
		time.Sleep(sleepTime) // exponential back-off

	}

	return fmt.Errorf("server %s failed to repond after %s", url, timeout)

}
