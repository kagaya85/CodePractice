package main

func hammingWeight(num uint32) (ans int) {
	for num != 0 {
		ans++
		num = (num - 1) & num
	}
	return
}
