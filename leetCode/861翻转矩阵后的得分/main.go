package main

func main() {

}

func matrixScore(A [][]int) int {
	m, n := len(A), len(A[0])

	ans := m * (1 << (n - 1))
	for j := 1; j < n; j++ {
		cnt := 0
		for _, row := range A {
			if row[j] == row[0] {
				cnt++
			}
		}
		if cnt < m-cnt {
			cnt = m - cnt
		}
		ans += cnt * (1 << (n - j - 1))
	}
	return ans
}
