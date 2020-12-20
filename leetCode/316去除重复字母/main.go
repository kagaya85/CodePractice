package main

func removeDuplicateLetters(s string) string {
	leftCnt := make(map[byte]int)
	for _, c := range []byte(s) {
		leftCnt[c]++
	}

	stack := []byte{}
	isInStack := make(map[byte]bool)

	for _, c := range []byte(s) {
		if !isInStack[c] {
			for len(stack) > 0 && c < stack[len(stack)-1] {
				lastCh := stack[len(stack)-1]
				if leftCnt[lastCh] == 0 {
					break
				}
				stack = stack[:len(stack)-1]
				isInStack[lastCh] = false
			}
			stack = append(stack, c)
			isInStack[c] = true
		}
		leftCnt[c]--
	}
	return string(stack)
}
