package main

import "fmt"

func main() {
	arr1 := [...]int{1, 2, 3, 4}
	arr2 := arr1
	fmt.Println("数组是copy=================")
	for k := range arr1 {
		fmt.Printf("%v ", &arr1[k])
	}
	fmt.Println("")
	for k := range arr2 {
		fmt.Printf("%v ", &arr2[k])
	}

	fmt.Println("")

	fmt.Println("\n切片是引用=================")
	slice1 := []int{1, 2, 3, 4}
	slice2 := slice1

	for k := range slice1 {
		fmt.Printf("%v ", &slice1[k])
	}
	fmt.Println("")
	for k := range slice2 {
		fmt.Printf("%v ", &slice2[k])
	}

	fmt.Println("")

	fmt.Printf("slice2的容量是:%d\n", cap(slice2))

	slice2 = append(slice2, 11, 12, 13, 14, 15, 16)

	fmt.Printf("slice2的容量是:%d\n", cap(slice2))

	fmt.Println(slice2)

	//fmt.Printf("arr type:%T\nslice type:%T\n", arr, slice)
	//fmt.Println(reflect.TypeOf(arr))
	//fmt.Println(reflect.TypeOf(slice))
}
