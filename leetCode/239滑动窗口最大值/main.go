package main

func maxSlidingWindow(nums []int, k int) []int {
	// 保存下标的优先队列
	q := []int{}
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			// 新入队列的元素比队尾元素大
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			// 窗口外的出队列
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}
