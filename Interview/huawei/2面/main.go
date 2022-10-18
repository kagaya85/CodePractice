package main

/*
输入字符串s和数字k，字符串s中都是数字，移除k个字符后，使得剩余的字符串的值是最大的。
输入描述:
数字字符串s，整数k。
输出描述:
字符串s1
示例1
输入
12345 2
输出
345
*/

import (
	"fmt"
)

func main() {
	s1, k1 := "12345", 2
	fmt.Println(foo(s1, k1))

	s2, k2 := "541322", 2
	fmt.Println(foo(s2, k2))

	s3, k3 := "541321", 2
	fmt.Println(foo(s3, k3))

	s4, k4 := "1150", 1
	fmt.Println(foo(s4, k4))

	s5, k5 := "53746", 2
	fmt.Println(foo(s5, k5))
}

func foo(s string, k int) string {
	arr := []byte(s)
	for k > 0 {
		flag := true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] < arr[i+1] {
				k--
				flag = false
				arr = append(arr[:i], arr[i+1:]...)
				break
			}
		}
		if flag {
			break
		}
	}

	for k > 0 {
		arr = arr[:len(arr)-1]
		k--
	}

	return string(arr)
}
