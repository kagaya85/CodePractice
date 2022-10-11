func isValid(s string) bool {
	stack := []byte{}

	for _, c := range []byte(s) {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			sl := len(stack)
			if sl == 0 {
				return false
			}
			switch c {
			case ')':
				if stack[sl-1] == '(' {
					stack = stack[:sl-1]
				} else {
					return false
				}
			case ']':
				if stack[sl-1] == '[' {
					stack = stack[:sl-1]
				} else {
					return false
				}
			case '}':
				if stack[sl-1] == '{' {
					stack = stack[:sl-1]
				} else {
					return false
				}
			default:
				return false
			}
		}
	}
	return len(stack) == 0
}