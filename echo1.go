package main

import (
	"fmt"
	"os"
)

func main() {
	step := " "
	var result string
	for i := 1; i < len(os.Args); i++ {
		result += step + os.Args[i]
	}
	fmt.Println(result)
}
