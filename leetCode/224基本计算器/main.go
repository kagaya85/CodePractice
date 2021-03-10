package main

func calculate(s string) (ans int) {
	var numS []int
	var signS []int
	sign := 1

	n := len(s)
	for i := 0; i < n; {
		c := s[i]
		switch c {
		case '(':
			numS = append(numS, ans)
			signS = append(signS, sign)
			ans = 0
			sign = 1
			i++
		case ')':
			sign = signS[len(signS)-1]
			signS = signS[:len(signS)-1]
			ans = numS[len(numS)-1] + sign*ans
			numS = numS[:len(numS)-1]
			i++
		case '+':
			sign = 1
			i++
		case '-':
			sign = -1
			i++
		case ' ':
			i++
			continue
		default:
			num := 0
			for ; i < n && s[i] >= '0' && s[i] <= '9'; i++ {
				num *= 10
				num += int(s[i] - '0')
			}
			ans += sign * num
		}
	}
	return
}
