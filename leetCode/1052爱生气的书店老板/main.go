package main

func maxSatisfied(customers []int, grumpy []int, X int) (ans int) {

	tot := 0
	for i, v := range customers {
		if grumpy[i] == 0 {
			tot += v
		}
	}

	left := 0
	for right, v := range customers {
		if right-left+1 > X {
			if grumpy[left] == 1 {
				tot -= customers[left]
			}
			left++
		}
		if grumpy[right] == 1 {
			tot += v
		}
		ans = max(ans, tot)
	}
	return ans
}

func max(arr ...int) int {
	var res = arr[0]
	for _, v := range arr[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
