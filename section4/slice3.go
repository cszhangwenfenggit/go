package main

import "fmt"

func main() {
	strings := []string{"one", "", "three"}

	//fmt.Printf("%q\n", noempty(strings))
	fmt.Println(noempty(strings))
	fmt.Println(strings)
}

func noempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}
