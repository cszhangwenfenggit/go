package main

import (
	"awesomeProject/section2/extend"
	"fmt"
	"log"
	"os"
	"time"
)

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("os.Getwd fail: %v", err)
	}
	log.Printf("Working directory = %s\n", cwd)
}

func main() {
	fmt.Println(extend.AbsoluteZeroC)

	fmt.Printf("100公斤=%.5f磅", extend.KToP(100))

	start := time.Now()
	fmt.Printf("x的种群统计=%d\n", extend.PopCount(98989898989))
	fmt.Printf("结束时间:%.5fs\n", time.Since(start).Seconds())

	x := "hello!"

	for i := 0; i < len(x); i++ {
		x := x[i]

		//fmt.Println(reflect.TypeOf(x))
		fmt.Println(x)

		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c\n", x)
		}
	}

	fmt.Println(uint8(33) == '!')

	//fmt.Printf("%c\n", 111+'A'-'a')

	var c int

	c = 200

	c <<= 2
	fmt.Printf("第 6行  - <<= 运算符实例，c 值为 = %d\n", c)

	c >>= 2
	fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c)

	fmt.Println(-5 % 3)
	fmt.Println(-5 % -3)
	fmt.Println(5.0 / 3.0)
	fmt.Println(5.0 / 4.0)

	a := 10.11
	fmt.Println(int(-a))

	str := "hello world"
	fmt.Println(str[4:])

	s := "left"
	t := s
	s += " right"

	fmt.Println(s, "\n", t)
}
