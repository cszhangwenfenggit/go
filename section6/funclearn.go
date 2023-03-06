package main

import "fmt"

type Student struct {
	Age  int
	Name string
}

func main() {
	student := Student{1, "Zhang"}

	learn(student)

	var student2 *Student = &Student{18, "Wen"}
	student2.tell() //等价 (*student2).tell()
}

func learn(s Student) {
	fmt.Println(s.Age)
}

func (s Student) tell() {
	fmt.Printf("Hello,My name is %s, I'm %d old!\n", s.Name, s.Age)
}
