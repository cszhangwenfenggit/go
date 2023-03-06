package main

import "fmt"

type ByTeCounter int

func (c *ByTeCounter) Write(p []byte) (int, error) {
	*c += ByTeCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByTeCounter
	c.Write([]byte("章文丰"))
	fmt.Println(c)

	c = 0
	var name string = "zhang"
	fmt.Fprintf(&c, "Hello %s", name)

	fmt.Println(c)
}
