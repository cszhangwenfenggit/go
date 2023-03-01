package main

import "fmt"

func main() {
	s := []string{"hello", "foo", "bar", "go", "abc", "zzz"}

	fmt.Println(reverse1(s))
	fmt.Println(reverse2(s))
	fmt.Println(s)
	reverse4(s)
	fmt.Println(s)

}

func reverse1(s []string) []string {
	reverseOfS := []string{}
	for i := range s {
		reverseOfS = append(reverseOfS, s[len(s)-i-1])
	}

	return reverseOfS
}

func reverse2(s []string) []string {
	reverseOfS := make([]string, len(s))
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		reverseOfS[i], reverseOfS[j] = s[j], s[i]
	}

	return reverseOfS
}

func reverse3(s []string) {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
}

func reverse4(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
