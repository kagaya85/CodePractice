package main

func candy(ratings []int) int {
	n := len(ratings)

	cur, dec, res, inc := 1, 0, 1, 1

	for i := 1; i < n; i++ {
		if ratings[i] >= ratings[i-1] {
			dec = 0
			if ratings[i] > ratings[i-1] {
				cur++
			} else {
				// 相等时从1开始分配
				cur = 1
			}
			inc = cur // 记录当前增长的最大值（同时也是递增长度）
			res += cur
		} else {
			dec++
			cur = 1
			if inc == dec {
				// 如果下降的长度等于上一次增长的长度，则把上一次增长的最后一个加入下降队列中
				dec++
			}
			res += dec
		}
	}

	return res
}
