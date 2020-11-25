package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aaaabbbbcccc"
	fmt.Println(sortString(s))
}

func sortString(s string) string {
	nums := make([]int, 26)
	var builder strings.Builder

	for _, c := range s {
		nums[c-'a']++
	}

	for builder.Len() != len(s) {
		for i := 0; i < 26; i++ {
			if nums[i] != 0 {
				nums[i]--
				builder.WriteByte('a' + byte(i))
			}
		}
		for i := 25; i >= 0; i-- {
			if nums[i] != 0 {
				nums[i]--
				builder.WriteByte('a' + byte(i))

			}
		}
	}

	return builder.String()
}
