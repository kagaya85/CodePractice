func candy(ratings []int) (res int) {
	n := len(ratings)
	left := make([]int, n)

	for i := range ratings {
		if i > 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}

	right := 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		res += max(left[i], right)
	}
	return
}

func max(arr ...int) (res int) {
	res = arr[0]
	for _, v := range arr[1:] {
		if v > res {
			res = v
		}
	}
	return
}