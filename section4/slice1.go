package main

import "fmt"

func main() {
	arr := [...]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
		13: "December",
	}

	test := arr[6:9]

	fmt.Println(arr)
	fmt.Println(test)
	//fmt.Println(test1)

	fmt.Println(len(test))
	fmt.Println(cap(test))
	fmt.Println(test[:cap(test)])
}
