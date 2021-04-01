package main

func clumsy(N int) (ans int) {
	stack := []int{N}
	sign := 0
	for n := N - 1; n > 0; n-- {
		switch sign % 4 {
		case 0: // *
			stack[len(stack)-1] *= n
		case 1: // /
			stack[len(stack)-1] /= n
		case 2: // +
			stack[len(stack)-1] += n
		default: // -
			stack = append(stack, -n)
		}
		sign++
	}

	for _, n := range stack {
		ans += n
	}
	return
}
