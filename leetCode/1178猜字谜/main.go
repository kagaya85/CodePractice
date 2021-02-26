package main

import "math/bits"

const puzzleLen = 7

func findNumOfValidWords(words []string, puzzles []string) []int {
	count := map[uint]int{}

	for _, w := range words {
		var mask uint
		for _, ch := range w {
			mask |= 1 << (ch - 'a')
		}
		if bits.OnesCount(mask) <= puzzleLen {
			count[mask]++
		}
	}

	ans := make([]int, len(puzzles))
	for i, p := range puzzles {
		first := 1 << (p[0] - 'a')
		mask := 0
		for _, c := range p[1:] {
			mask |= 1 << (c - 'a')
		}
		subset := mask
		for {
			ans[i] += count[uint(subset|first)] // 0 也是合法子集之一
			subset = (subset - 1) & mask        // On 枚举所有子集
			if subset == mask {
				break
			}
		}
	}

	return ans
}
