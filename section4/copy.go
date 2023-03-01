// copy 解析
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("原始值：", slice)

	sc := []int{11, 12, 13}
	fmt.Println("覆盖值：", sc)

	copy(slice[3:], sc) //sc覆盖slice目标数据段([3:])

	fmt.Println("修改值：", slice)
	fmt.Println("变化值：", slice[3:(3+len(sc))])
}
