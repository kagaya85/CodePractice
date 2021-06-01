package main

func canEat(candiesCount []int, queries [][]int) []bool {
	n := len(candiesCount)
	presum := make([]int, n+1)
	for i := range candiesCount {
		presum[i+1] = presum[i] + candiesCount[i]
	}

	ans := make([]bool, len(queries))

	for i, q := range queries {
		favtype, favday, cap := q[0], q[1], q[2]
		eatmin, eatmax := favday+1, (favday+1)*cap
		candymin, candymax := presum[favtype]+1, presum[favtype]+candiesCount[favtype]
		ans[i] = eatmin <= candymax && eatmax >= candymin
	}

	return ans
}
