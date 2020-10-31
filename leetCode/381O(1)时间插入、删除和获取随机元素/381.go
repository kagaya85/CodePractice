package main

import "math/rand"

func main() {

}

// RandomizedCollection sturct
type RandomizedCollection struct {
	index map[int]map[int]struct{}
	nums  []int
}

// Constructor  Initialize your data structure here.
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		index: map[int]map[int]struct{}{},
		nums:  []int{},
	}
}

// Insert  Inserts a value to the collection. Returns true if the collection did not already contain the specified element.
func (rc *RandomizedCollection) Insert(val int) bool {
	iset, ok := rc.index[val]

	if !ok {
		iset = map[int]struct{}{}
		rc.index[val] = iset
	}
	iset[len(rc.nums)] = struct{}{}
	rc.nums = append(rc.nums, val)

	return !ok
}

// Remove a value from the collection. Returns true if the collection contained the specified element.
func (rc *RandomizedCollection) Remove(val int) bool {
	iset, ok := rc.index[val]
	isetlen := len(iset)
	if !ok || isetlen == 0 {
		return false
	}

	// 获取删除数字的idx
	var idx int
	for k := range iset {
		idx = k
		break
	}
	delete(iset, idx)

	// 交换最后一个数字，删除
	numslen := len(rc.nums)
	rc.nums[idx] = rc.nums[numslen-1]
	rc.nums = rc.nums[:numslen-1]

	// 如果删除的不是最后一个数，修改被交换数的index
	if idx < numslen-1 {
		delete(rc.index[rc.nums[idx]], numslen-1)
		rc.index[rc.nums[idx]][idx] = struct{}{}
	}

	return true
}

// GetRandom Get a random element from the collection.
func (rc *RandomizedCollection) GetRandom() int {
	return rc.nums[rand.Intn(len(rc.nums))]
}
