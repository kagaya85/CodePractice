func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i, a := range nums {
		if i > 0 && a == nums[i-1] {
			continue
		}

		j, k := i+1, n-1
		for j < k {
			b, c := nums[j], nums[k]
			if j > i+1 && b == nums[j-1] {
				j++
				continue
			}
			if k < n-1 && c == nums[k+1] {
				k--
				continue
			}

			if b+c == -a {
				ans = append(ans, []int{a, b, c})
				j++
				k--
			} else if b+c > -a {
				k--
			} else {
				j++
			}
		}
	}
	return
}