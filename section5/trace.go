package main

import (
	"fmt"
	"log"
	"time"
)

// !+main
func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses

	fmt.Println("Do something")
	// ...lots of work...
	time.Sleep(3 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

//!-main

func main() {
	bigSlowOperation()
}
