package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	sum := 1
	for sum < 3 {
		sum += sum
	}
	fmt.Println(sum)

	if v := rand.Intn(10); v < 5 {
		fmt.Printf("随机值 %d < 5\n", v)
	} else {
		fmt.Printf("随机值 %d >= 5\n", v)
	}

	fmt.Println(runtime.GOOS)

	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In Two Days")
	default:
		fmt.Println("Too far away!")
		//
	}

	now := time.Now().Hour()
	switch {
	case now < 12:
		fmt.Println("Good Morning!")
	case now < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	defer fmt.Println("I'm defer")
	fmt.Println("I'm running before defer")
}
