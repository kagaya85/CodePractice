package main

func main() {

}

func maxNumber(nums1 []int, nums2 []int, k int) (res []int) {

	begin := 0
	if k > len(nums2) {
		begin = k - len(nums2)
	}

	for i := begin; i <= len(nums1) && i <= k; i++ {
		sub1 := maxSub(nums1, i)
		sub2 := maxSub(nums2, k-i)
		newNums := merge(sub1, sub2)
		if less(res, newNums) {
			res = newNums
		}
	}

	return
}

// 维护单调栈
func maxSub(nums []int, n int) (sub []int) {
	for i, v := range nums {
		l := len(sub)
		for l > 0 && v > sub[l-1] && l+len(nums)-i > n {
			sub = sub[:l-1] // pop
			l--
		}
		if l < n {
			sub = append(sub, v)
		}
	}
	return
}

func merge(a, b []int) []int {
	nums := make([]int, len(a)+len(b))
	for i := range nums {
		if less(a, b) {
			nums[i], b = b[0], b[1:]
		} else {
			nums[i], a = a[0], a[1:]
		}
	}
	return nums
}

func less(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}
