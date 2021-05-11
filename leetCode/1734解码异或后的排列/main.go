package main

func decode(encoded []int) []int {
	n := len(encoded)
	total := 0
	for i := 1; i <= n+1; i++ {
		total ^= i
	}
	rest := 0
	for i := 1; i < n; i += 2 {
		rest ^= encoded[i]
	}
	ans := make([]int, n+1)
	ans[0] = total ^ rest
	for i, v := range encoded {
		ans[i+1] = ans[i] ^ v
	}
	return ans
}
