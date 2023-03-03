package main

import "fmt"

//这里还要提醒一句,对于[goroutin(程道)],[切片],[映射]这三种类型来说,只有形参,而且不需要加[*]号.

//另外,对于参数类型是[interface]的函数参数,只有实参,而且不会将[interface]结构所包含的地址复制!

type coder interface {
	code()
	//debug()
}

type Gopher struct {
	language string
}

// 实现了接收者是值类型的方法(p)，会隐含地也实现了接收者是指针类型的方法(*p)
func (p Gopher) code() {
	fmt.Printf("I am coding %s language\n", p.language)
}

// 实现了接收者是指针类型的方法(*p)，不会自动生成对应接收者是值类型的方法(p)
func (p *Gopher) debug() {
	fmt.Printf("I am debuging %s language\n", p.language)
}

func main() {
	var p coder = Gopher{"Go"}
	p.code()
	//p.debug()

	//var t = &Gopher{"Go"}
	//t.code()
	//t.debug()
}
