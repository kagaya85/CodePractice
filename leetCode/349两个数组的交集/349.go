package main

func main() {

}

func intersection(nums1, nums2 []int) (ans []int) {

	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})

	for _, num := range nums1 {
		if _, has := set1[num]; !has {
			set1[num] = struct{}{}
		}
	}

	for _, num := range nums2 {
		_, has1 := set1[num]
		_, has2 := set2[num]
		if has1 && !has2 {
			ans = append(ans, num)
			set2[num] = struct{}{}
		}
	}

	return
}
