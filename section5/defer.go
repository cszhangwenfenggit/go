package main

import "fmt"

func main() {
	a := a()
	fmt.Println("defer a:", a)

	c := c()
	fmt.Println("defer c:", c)
}

func a() int {
	i := 0
	//fmt.Printf("%d", i)
	defer fmt.Println("defer a get:", i)
	i++
	return i
}

func c() (i int) {
	defer func() {
		fmt.Println("defer c get:", i)
		i++
	}()
	return 1
}
