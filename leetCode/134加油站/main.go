package main

import "fmt"

func main() {

	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}

	fmt.Println(canCompleteCircuit(gas, cost))
}

func canCompleteCircuit(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		gasSum, costSum, offset := 0, 0, 0
		for offset < n {
			j := (i + offset) % n
			gasSum += gas[j]
			costSum += cost[j]
			if costSum > gasSum {
				break
			}
			offset++
		}
		if offset == n {
			return i
		}
		i += offset + 1
	}
	return -1
}
