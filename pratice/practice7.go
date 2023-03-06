package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

type IPAddr [4]byte

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func (ip IPAddr) String() string {
	var newIp []string
	for _, v := range ip {
		newIp = append(newIp, strconv.Itoa(int(v)))
	}

	return fmt.Sprintf("%s", strings.Join(newIp, ","))
}

func main() {
	p1 := Person{"p1", 18}
	p2 := Person{"p2", 36}

	fmt.Println(p1)
	fmt.Println(p2)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
