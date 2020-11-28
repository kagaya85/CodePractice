package main

import "fmt"

func main() {
	arr := []int{1, 3, 2, 3, 2, 1}
	fmt.Println(reversePairs(arr))
}

func reversePairs(nums []int) (ans int) {

	n := len(nums)

	if n <= 1 {
		return 0
	}

	a := append([]int(nil), nums[:n>>1]...)
	b := append([]int(nil), nums[n>>1:]...)

	ans = reversePairs(a) + reversePairs(b)

	idx := 0
	for _, v := range a {
		for idx < len(b) && v > b[idx]<<1 {
			idx++
		}
		ans += idx
	}

	p1, p2 := 0, 0
	for i := range nums {
		if p1 < len(a) && (p2 >= len(b) || a[p1] <= b[p2]) {
			nums[i] = a[p1]
			p1++
		} else {
			nums[i] = b[p2]
			p2++
		}
	}

	return ans
}
