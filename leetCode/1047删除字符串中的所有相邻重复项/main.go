package main

func removeDuplicates(S string) string {
	var stack []byte
	for _, c := range []byte(S) {
		if len(stack) == 0 || stack[len(stack)-1] != c {
			stack = append(stack, c)
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return string(stack)
}
