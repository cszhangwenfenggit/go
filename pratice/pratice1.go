package main

import (
	"fmt"
	"pratice/other"
)

func main() {
	for i := 0; i < 100; i++ {
		//fmt.Println(rand.Intn(10) + 15)
	}

	other.Pizza("100")

	fmt.Println(other.Add(100, 10))

	swapX, swapY := other.Swap("Hello", "World")
	fmt.Println(swapX, swapY)

	fmt.Println(other.ReturnP(10))
}
