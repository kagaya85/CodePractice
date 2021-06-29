package main

func convertToTitle(colNum int) string {
	ans := []byte{}
	for colNum > 0 {
		colNum--
		ans = append(ans, 'A'+byte(colNum%26))
		colNum /= 26
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-i-1] = ans[n-i-1], ans[i]
	}
	return string(ans)
}
