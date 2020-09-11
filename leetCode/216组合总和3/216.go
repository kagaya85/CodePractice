package main

import (
	"fmt"
)

func main() {
	
	k := 3
	n := 9
	
	ans := combinationSum3(k, n)
	
	fmt.Println(ans)
}

func combinationSum3(k int, n int) (ans [][]int) {
	
	var seq []int
	var dfs func(start, num, rest int)
	
	dfs = func(start, num, rest int) {
		
//		fmt.Printf("seq: %v, start: %d, num: %d, rest: %d\n", seq, start, num, rest)
		
		if rest == 0 && num == 0 {
			ans = append(ans, append([]int(nil), seq...))
			return
		}
		if (start >= 10 || rest < start || num <= 0) {
			return
		}
		
		for i := start; i < 10; i++ {
			seq = append(seq, i)
			dfs(i+1, num-1, rest-i)
			seq = seq[:len(seq)-1]
			if rest-i < 0 {
				break
			}
		}		
	}
	
	dfs(1, k, n)
	return
}