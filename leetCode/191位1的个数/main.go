package main

func hammingWeight(num uint32) (cnt int) {
	for num > 0 {
		num &= num - 1
		cnt++
	}
	return
}
