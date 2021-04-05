package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	ans := make([]int, 0, m+n)
	i, j := 0, 0

	for i < m && j < n {
		if nums1[i] < nums2[j] {
			ans = append(ans, nums1[i])
			i++
		} else {
			ans = append(ans, nums2[j])
			j++
		}
	}

	if i == m {
		ans = append(ans, nums2[j:]...)
	}

	if j == n {
		ans = append(ans, nums1[i:]...)
	}

	copy(nums1, ans)
}
