package main

func main() {

}

func totalNQueens(n int) (ans int) {

	var dfs func(int, int, int, int)
	dfs = func(row, col, leftd, rightd int) {
		if row >= n {
			ans++
			return
		}

		// 将01反转，并清空n位以前的1，反转后1所在的位置可以摆放皇后
		bits := ^(col | leftd | rightd) & ((1 << n) - 1)
		for bits > 0 {
			pick := bits & -bits // 与自己的补码做与操作，获取最后一位1（可放皇后的位置）
			dfs(row+1, col|pick, (leftd|pick)<<1, (rightd|pick)>>1)
			bits &= bits - 1 // 与自己-1做与操作，将最后一位1清零
		}
	}

	dfs(0, 0, 0, 0)

	return
}
