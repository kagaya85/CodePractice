package main

type NumArray struct {
	presum []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	presum := make([]int, n+1)
	for i, v := range nums {
		presum[i+1] = presum[i] + v
	}
	return NumArray{presum}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.presum[j+1] - this.presum[i]
}
