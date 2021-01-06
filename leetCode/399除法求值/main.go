package main

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 给变量编号
	id := map[string]int{}
	for _, eq := range equations {
		a, b := eq[0], eq[1]
		if _, has := id[a]; !has {
			id[a] = len(id)
		}
		if _, has := id[b]; !has {
			id[b] = len(id)
		}
	}

	// 并查集
	f := make([]int, len(id))
	// 权值（父子比值）集合
	w := make([]float64, len(id))
	for i := range f {
		f[i] = i
		w[i] = 1
	}

	var find func(int) int
	find = func(x int) int {
		if x != f[x] {
			father := find(f[x])
			w[x] *= w[f[x]]
			f[x] = father
		}
		return f[x]
	}

	union := func(from, to int, val float64) {
		fFrom, fTo := find(from), find(to)
		w[fFrom] = val * w[to] / w[from]
		f[fFrom] = fTo
	}

	for i, eq := range equations {
		union(id[eq[0]], id[eq[1]], values[i])
	}

	ans := make([]float64, len(queries))
	for i, q := range queries {
		start, hasS := id[q[0]]
		end, hasE := id[q[1]]
		if hasS && hasE && find(start) == find(end) {
			ans[i] = w[start] / w[end]
		} else {
			ans[i] = -1.0
		}
	}

	return ans
}
