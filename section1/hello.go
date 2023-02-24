package main

import "fmt"

func main() {
	data := make(map[string]map[string]int)

	data["text.txt"] = map[string]int{"name": 213}

	//fmt.Println(data)
	fmt.Println(data["text.txt"]["name"])
}
