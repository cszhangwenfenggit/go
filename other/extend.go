package other

import "fmt"

func Pizza(size string) {
	fmt.Printf("Pizza size is %s\n", size)
}

func Add(x, y int) int {
	return x + y
}

func Swap(x, y string) (string, string) {
	return x, y
}

func ReturnP(val int) (x, y int) {
	x = val*10 - 9
	y = val + x
	return
}
