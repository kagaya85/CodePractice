package main

func subarraysWithKDistinct(A []int, K int) (ans int) {
	n := len(A)
	count1, count2 := make(map[int]int, n+1), make(map[int]int, n+1)
	left1, left2 := 0, 0
	diff1, diff2 := 0, 0
	for _, v := range A {
		if _, ok := count1[v]; !ok {
			diff1++
		}
		count1[v]++
		for diff1 > K {
			num := A[left1]
			count1[num]--
			if count1[num] == 0 {
				delete(count1, num)
				diff1--
			}
			left1++
		}
		if _, ok := count2[v]; !ok {
			diff2++
		}
		count2[v]++
		for diff2 > K-1 {
			num := A[left2]
			count2[num]--
			if count2[num] == 0 {
				delete(count2, num)
				diff2--
			}
			left2++
		}
		ans += left2 - left1
	}
	return
}
