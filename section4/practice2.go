package main

import "fmt"

func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(reverse(arr))
}

func reverse(a [6]int) [6]int {
	for i := 0; i < len(a)/2; i++ {
		j := len(a) - 1 - i
		a[i], a[j] = a[j], a[i]
	}

	return a
}
