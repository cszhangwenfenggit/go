package main

import "fmt"

func main() {
	slice := []int{1, 2, 4, 5, 7, 9}

	fmt.Println(remove(slice, 1))

	//fmt.Println(len(slice))
	//copy(slice[3:], slice[5+1:])
	//
	//fmt.Println(slice[:len(slice)-1])

	//slice1 := []int{1, 2, 4, 5, 7, 9}
	//
	//slice2 := []int{11, 12, 13}
	//
	////copy(slice1, slice2)
	//copy(slice2, slice1)
	//fmt.Println(slice2)
}

func remove(slice []int, target int) []int {
	copy(slice[target:], slice[target+1:])

	return slice[:len(slice)-1]
}
