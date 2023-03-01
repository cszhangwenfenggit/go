package main

import "fmt"

type Employee struct {
	ID       int
	Name     string
	Address  string
	Position string
	Salary   int
}

type Point struct{ x, y int }

func main() {

	pp := &Point{1, 2}
	zz := new(Point)
	*zz = Point{1, 2}
	fmt.Println(pp)
	fmt.Println(zz)
	fmt.Printf("%v\n", &pp.x)
	fmt.Printf("%v\n", &zz.x)

	var jack Employee

	jack.ID = 1

	name := &jack.Name

	*name = "ZhangWenFeng"

	fmt.Println(jack.ID, jack.Name)

	var employeeOfTheMonth *Employee = &jack

	employeeOfTheMonth.Position = "World"

	fmt.Println(employeeOfTheMonth.Position)
	//fmt.Printf("%v, %v", &jack, &employeeOfTheMonth)
}

//func getEmployeeById(id int) *Employee {
//
//}
