package main

import "fmt"

//练习4.5 编写一个函数用于去除string slice相邻重复字符串的元素
//findindex可以找到需要除去的元素的下标
// func findindex(s []string) []int {
// 	var x []int
// 	for i := 0; i < len(s)-1; i++ {
// 		if s[i] == s[i+1] {
// 			x = append(x, i+1)
// 		}
// 	}
// 	return x
// }

// clean 用于除去元素
func clean(s []string) []string {
	var x []int //x切片用于记录需要删除字符串的下标
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			x = append(x, i+1)
		}
	}
	for t, v := range x {
		copy(s[v-t:], s[v+1-t:]) //每覆盖一个前面字符串，下标集体减一，覆盖t次下标减少t
	}
	fmt.Println(x)
	return s[:len(s)-len(x)]
}

func main() {
	s := []string{"1", "1", "2", "3", "44", "44", "44", "44", "5", "5", "6241", "6241", "124142", "7788", "asd", "asd", "7788"}
	fmt.Println(s)
	s = clean(s)
	fmt.Println(s)
}
