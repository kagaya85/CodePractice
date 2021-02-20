package main

type entry struct {
	count, l, r int
}

func findShortestSubArray(nums []int) (ans int) {
	cnt := make(map[int]*entry)
	for i, v := range nums {
		if c, ok := cnt[v]; !ok {
			cnt[v] = &entry{1, i, i}
		} else {
			c.count++
			c.r = i
		}
	}

	maxCnt := 0
	for _, v := range cnt {
		if v.count > maxCnt {
			maxCnt, ans = v.count, v.r-v.l+1
		} else if v.count == maxCnt {
			ans = min(ans, v.r-v.l+1)
		}
	}
	return
}

func min(arr ...int) int {
	res := arr[0]
	for _, v := range arr[1:] {
		if v < res {
			res = v
		}
	}
	return res
}
