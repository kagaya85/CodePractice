package main

func hammingDistance(x int, y int) (ans int) {
	// return bits.OnesCount(uint(x ^ y))
	for s := x ^ y; s > 0; s &= s - 1 {
		ans++
	}
	return ans
}
