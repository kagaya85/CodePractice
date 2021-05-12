package main

func xorQueries(arr []int, queries [][]int) []int {
	n := len(queries)
	prefix := make([]int, len(arr)+1)

	for i, v := range arr {
		prefix[i+1] = prefix[i] ^ v
	}

	ans := make([]int, n)
	for i, q := range queries {
		from, to := q[0], q[1]
		ans[i] = prefix[from] ^ prefix[to+1]
	}

	return ans
}
