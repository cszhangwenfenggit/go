package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//words := []int{0, 1}
	//for _, value := range words {
	//	//fmt.Println(i, value)
	//	words = append(words, value)
	//	//value += 10
	//}
	//fmt.Printf("words的值是%+v\n", words)
	counts := make(map[string]map[string]int)
	for _, fileName := range os.Args[1:] {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup open error%v\n", err)
			continue
		}
		counts[fileName] = map[string]int{}
		for _, line := range strings.Split(string(data), "\n") {
			counts[fileName][line]++
		}
	}
	for file, data := range counts {
		for line, n := range data {
			if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", file, n, line)
			}
		}
	}

}
