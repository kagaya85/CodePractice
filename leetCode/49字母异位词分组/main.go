package main

import "sort"

func main() {

}

func groupAnagrams(strs []string) [][]string {
	count := map[string][]string{}

	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		count[sortedStr] = append(count[sortedStr], str)
	}
	ans := make([][]string, 0, len(count))
	for _, v := range count {
		ans = append(ans, v)
	}
	return ans
}
