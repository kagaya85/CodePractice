package main

import "fmt"

func main() {
	input := []int{1, 0, 1, 0, 0}
	fmt.Println(numRabbits(input))
}

func numRabbits(answers []int) (ans int) {
	// 报数不同的兔子颜色肯定不同，
	count := make(map[int]int)
	for _, n := range answers {
		count[n]++
	}

	for k, v := range count {
		numPerColor := k + 1
		// 相同报数的兔子数量 / 对应一种颜色兔子最少数量 （向上取整）= 最少的不同颜色数量
		// ans += int(math.Ceil(float64(v)/float64(numPerColor))) * numPerColor
		ans += ((k + v) / numPerColor) * numPerColor
	}
	return
}
