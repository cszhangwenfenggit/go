package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	fmt.Println(os.Stdin)

	fmt.Println("命令名称:", os.Args[0])

	fmt.Println(time.Now())
	for i := 1; i < len(os.Args); i++ {
		fmt.Println("索引:", i, ";值:", os.Args[i])
	}
	fmt.Println(time.Now())
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Now())
}
