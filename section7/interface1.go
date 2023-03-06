package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof~~"
}

type Cat struct{}

func (c *Cat) Speak() string {
	return "Mio~~"
}

type PhpProgrammer struct{}

func (p PhpProgrammer) Speak() string {
	return "Hello~~"
}

func main() {
	animals := []Animal{Dog{}, new(Cat), PhpProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

	//var t Animal = &Cat{}
}
