package main

import "fmt"

func main() {
	arr := []string{"a", "b", "c", "d", "a", "c", "t"}

	//arrx := []string

	for i, v := range arr[3:] {
		fmt.Printf("i:%d, v:%v\n", i, v)
	}

	copy(arr[3:], arr[4:])
	fmt.Println(arr)

	//arr1 := arr[:len(arr)-1]
	fmt.Println(arr[1:])
	fmt.Println(arr[2:])

}
