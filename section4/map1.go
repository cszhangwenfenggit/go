package main

import "fmt"

func main() {
	map1 := make(map[string]int)

	map1["name"] = 21

	fmt.Println(map1)

	var map2 map[string]int
	//fmt.Println(map2 == nil)
	map2 = map[string]int{"name": 123}
	map2["sex"] = 1
	fmt.Println(map2)
	//fmt.Println(map1 == nil)
	//fmt.Println(map2 == nil)
	//fmt.Printf("%v", map1)

	a, ok := map2["age"]
	fmt.Println(a, ok)
}
