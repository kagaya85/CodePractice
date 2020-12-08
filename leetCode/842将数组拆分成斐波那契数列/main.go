package main

import "math"

func main() {

}

func splitIntoFibonacci(S string) (res []int) {
	n := len(S)

	var backtrack func(int, int, int) bool
	backtrack = func(index, pre, sum int) bool {
		if index == n {
			return len(res) >= 3
		}

		cur := 0
		for i := index; i < n; i++ {
			if i > index && S[index] == '0' {
				return false
			}

			cur = cur*10 + int(S[i]-'0')
			if cur > math.MaxInt32 {
				return false
			}

			if len(res) >= 2 {
				if cur < sum {
					continue
				}
				if cur > sum {
					return false
				}
			}

			res = append(res, cur)
			if backtrack(i+1, cur, pre+cur) {
				return true
			}
			res = res[:len(res)-1]
		}
		return false
	}

	backtrack(0, 0, 0)
	return
}
