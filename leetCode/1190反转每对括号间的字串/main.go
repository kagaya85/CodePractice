package main

func reverseParentheses(s string) string {
	n := len(s)
	pairidx := map[int]int{}
	stack := []int{}
	for i, ch := range s {
		if ch == '(' {
			stack = append(stack, i)
		} else if ch == ')' {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pairidx[i], pairidx[j] = j, i
		}
	}

	ans := []byte{}
	for i, direct := 0, 1; i < n; i += direct {
		if s[i] == '(' || s[i] == ')' {
			i = pairidx[i]
			direct = -direct
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}
