package main

import "fmt"

type Student struct {
	name string

	school string
}

func main() {
	//s := &Student{name: "Ben", school: "JHL"}
	s := Student{"Hello", "World"}

	fmt.Println(s)
	fmt.Println(s.name)
	fmt.Println(s.school)

	var a [10]string

	a[0] = "name"
	a[9] = "name"

	fmt.Println(a)
	fmt.Println(a[0])

	slice := []int{1, 2, 3, 4}
	fmt.Println(slice[:2])

	slice1 := []struct {
		i int
		b bool
	}{
		{1, false},
		{2, false},
	}

	fmt.Println(slice1)

	//var slice2 []int
	var slice2 = make([]int, 3)
	var slice3 = make([]int, 3)

	fmt.Println(len(slice2), cap(slice2))
	fmt.Println(cap(slice3))

	m := make(map[string]string)

	m["name"] = "zhang"
	fmt.Println(m["name"])
	m["name"] = "wen"
	fmt.Println(m["name"])

	delete(m, "name")
	fmt.Println(m["name"])

	v, ok := m["name"]
	fmt.Printf("v:%s ok:%t\n", v, ok)

	str := "Hello World"

	fmt.Println(str[0:4])
}
