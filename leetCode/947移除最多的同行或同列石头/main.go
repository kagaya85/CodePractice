package main

func removeStones(stones [][]int) int {
	// 使用Map代替底层数组
	ancestor := make(map[int]int)
	// 连通分量计数
	count := 0

	var find func(int) int
	find = func(x int) int {
		if a, has := ancestor[x]; !has {
			ancestor[x] = x
			// 连通分量+1
			count++
		} else if a != x {
			ancestor[x] = find(ancestor[x])
		}
		return ancestor[x]
	}

	union := func(from, to int) {
		fa, ta := find(from), find(to)
		if fa == ta {
			// 已经是一个连通分量
			return
		}

		ancestor[fa] = ta
		// 连通分量-1
		count--
	}

	for _, stone := range stones {
		// 将横坐标投射到另一段区域来区分横纵坐标，注意不能让横纵坐标的值域出现重叠（0值）
		union(-stone[0]-1, stone[1])
	}

	return len(stones) - count
}
