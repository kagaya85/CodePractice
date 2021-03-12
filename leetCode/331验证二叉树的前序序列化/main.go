package main

func isValidSerialization(preorder string) bool {
	n := len(preorder)
	slots := 1
	for i := 0; i < n; {
		if slots == 0 {
			return false
		}
		if preorder[i] == ',' {
			i++
		} else if preorder[i] == '#' {
			slots--
			i++
		} else {
			for i < n && preorder[i] != ',' {
				i++
			}
			slots++ // slots = slots - 1 + 2
		}
	}
	return slots == 0

}
