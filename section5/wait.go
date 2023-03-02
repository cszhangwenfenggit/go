package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "http://www.ba212idu123.com"
	if err := WaitForServer(url); err != nil {
		log.SetPrefix("wait:")
		log.Fatalf("website down: %v\n", err)
	}
}

func WaitForServer(url string) error {
	const timeout = 10 * time.Second
	deadline := time.Now().Add(timeout)
	for t := 0; time.Now().Before(deadline); t++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		fmt.Printf("server not responing (%s), retrying(%d)...\n", err, t)
		time.Sleep(time.Second << uint(t))
	}

	return fmt.Errorf("server %s failed to response after %s", url, timeout)
}
