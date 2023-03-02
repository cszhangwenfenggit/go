package main

import "fmt"

func main() {
	f := square()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func square() func() int {
	var x int
	fmt.Println("square 初始化")
	return func() int {
		x++
		return x * x
	}
}
