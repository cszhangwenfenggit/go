package main

import "fmt"

func main() {
	//todo 删除切片相邻的重复字符串元素
	slice := []string{"a", "v", "s", "f", "f", "f", "r", "d", "a", "a", "o", "q"}
	//slice := []string{"a", "v", "s", "f", "r", "d", "a", "o", "q"}

	fmt.Println(slice)

	//fmt.Println(removeRepeat(slice))
	fmt.Println(removeRepeat2(slice))
	//removeRepeat2(slice)
	//fmt.Println(slice)
	//fmt.Println(removeRepeat(slice))
}

func removeRepeat(s []string) []string {
	sln := len(s)
	//fmt.Println(s)
	//fmt.Println("初始化", sln)
	for i := 0; i < sln; i++ {
		//fmt.Println("循环", i)
		if i != sln-1 && s[i] == s[i+1] {
			//fmt.Println("重置前", sln)
			copy(s[i:], s[i+1:])

			s = s[:len(s)-1]

			//fmt.Println("重置后", len(s))

			return removeRepeat(s)
		}
	}

	return s
	//return s[:len(s)-1]
	//for i, str := range s {
	//	if i <= len(s)-1;

	//fmt.Println(str)
}

// clean 用于除去元素
func removeRepeat2(s []string) []string {
	var x []int //x切片用于记录需要删除字符串的下标
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			x = append(x, i+1)
		}
	}

	fmt.Println(x)

	for t, v := range x {
		fmt.Printf("v(%d)-t(%d)=%d\n", v, t, v-t)

		copy(s[v-t:], s[v+1-t:]) //每覆盖一个前面字符串，下标集体减一，覆盖t次下标减少t
		fmt.Println(s)
	}

	return s[:len(s)-len(x)]
}
