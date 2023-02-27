package main

import (
	"flag"
	"fmt"
)

func main() {
	boolArgPtr := flag.Bool("arg", false, "This is a bool argument")
	flag.Parse()
	fmt.Println("Bool Arg", *boolArgPtr)
}
