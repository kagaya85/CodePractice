package main

import "fmt"

func main() {
	s := "aaaabbbbcccc"
	fmt.Println(sortString(s))
}

func sortString(s string) string {
	nums := make([]int, 26)
	ans := []byte{}

	for _, c := range s {
		nums[c-'a']++
	}

	for len(ans) != len(s) {
		for i := 0; i < 26; i++ {
			if nums[i] != 0 {
				nums[i]--
				ans = append(ans, 'a'+byte(i))
			}
		}
		for i := 25; i >= 0; i-- {
			if nums[i] != 0 {
				nums[i]--
				ans = append(ans, 'a'+byte(i))
			}
		}
	}

	return string(ans)
}
