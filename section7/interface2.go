package main

import "fmt"

func PrintAll(vals []interface{}) {
	for _, i := range vals {
		fmt.Println(i)
	}
}

func main() {
	strings := []string{"Hello", "World"}

	vals := make([]interface{}, len(strings))
	for i, v := range strings {
		vals[i] = v
	}

	PrintAll(vals)
}
