package main

const high = 30

type trie struct {
	Left, Right *trie
}

func (t *trie) add(num int) {
	cur := t
	for i := high; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if cur.Left == nil {
				cur.Left = &trie{}
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = &trie{}
			}
			cur = cur.Right
		}
	}
}

func (t *trie) check(num int) (ans int) {
	cur := t
	for i := high; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if cur.Right != nil {
				cur = cur.Right
				ans = ans<<1 + 1
			} else {
				cur = cur.Left
				ans <<= 1
			}
		} else {
			if cur.Left != nil {
				cur = cur.Left
				ans = ans<<1 + 1
			} else {
				cur = cur.Right
				ans <<= 1
			}
		}
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaximumXOR(nums []int) (ans int) {
	root := &trie{}

	for i := 1; i < len(nums); i++ {
		root.add(nums[i-1])
		ans = max(ans, root.check(nums[i]))
	}

	return
}
