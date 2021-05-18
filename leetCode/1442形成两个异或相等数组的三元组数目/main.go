package main

func countTriplets(arr []int) (ans int) {
	sameCnt := map[int]int{}
	indexSum := map[int]int{}
	pre := 0
	for i, val := range arr {
		cur := pre ^ val
		if cnt, has := sameCnt[cur]; has {
			ans += cnt*i - indexSum[cur]
		}
		sameCnt[pre]++
		indexSum[pre] += i
		pre = cur
	}
	return
}
