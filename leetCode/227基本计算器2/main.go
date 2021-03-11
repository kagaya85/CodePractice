package main

import "sort"

const (
	plus = iota
	minus
	multi
	divid
	left
	right
)

type stack struct {
	sort.IntSlice
}

func (s *stack) Pop() int {
	e := s.IntSlice[s.Len()-1]
	s.IntSlice = s.IntSlice[:s.Len()-1]
	return e
}

func (s *stack) Push(e int) {
	s.IntSlice = append(s.IntSlice, e)
}

func (s *stack) Peek() int {
	return s.IntSlice[s.Len()-1]
}

func calculate(s string) int {
	var numS stack
	var signS stack

	cal := func() int {
		if numS.Len() < 2 {
			panic(numS.Len())
		}
		sign := signS.Pop()
		num2, num1 := numS.Pop(), numS.Pop()
		switch sign {
		case plus:
			return num1 + num2
		case minus:
			return num1 - num2
		case multi:
			return num1 * num2
		case divid:
			return num1 / num2
		}
		panic("error")
	}

	n := len(s)
	for i := 0; i < n; {
		c := s[i]
		switch c {
		case '(':
			signS.Push(left)
			i++
		case ')':
			numS.Push(cal())
			signS.Pop() // pop '('
			i++
		case '+':
			if signS.Len() > 0 && signS.Peek() != '(' && numS.Len() > 1 {
				numS.Push(cal())
			}
			signS.Push(plus)
			i++
		case '-':
			if signS.Len() > 0 && signS.Peek() != '(' && numS.Len() > 1 {
				numS.Push(cal())
			}
			signS.Push(minus)
			i++
		case '*':
			signS.Push(multi)
			i++
		case '/':
			signS.Push(divid)
			i++
		case ' ':
			i++
		default:
			num := 0
			for ; i < n && s[i] >= '0' && s[i] <= '9'; i++ {
				num *= 10
				num += int(s[i] - '0')
			}
			numS.Push(num)
			if signS.Len() > 0 {
				if preSign := signS.Peek(); preSign == multi || preSign == divid {
					numS.Push(cal())
				}
			}
		}
	}
	for signS.Len() > 0 {
		numS.Push(cal())
	}
	return numS.Pop()
}
