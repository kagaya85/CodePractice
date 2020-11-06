package main

import (
	"fmt"
	"sort"
)

// SortByBits sort by bits
type SortByBits []int

func (a SortByBits) Len() int      { return len(a) }
func (a SortByBits) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByBits) Less(i, j int) bool {
	counti := countBit1(a[i])
	countj := countBit1(a[j])

	if counti == countj {
		return a[i] < a[j]
	}
	return counti < countj
}

func main() {
	arr := []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}
	fmt.Println(sortByBits(arr))
}

func sortByBits(arr []int) []int {

	sort.Sort(SortByBits(arr))

	return arr
}

func countBit1(num int) int {
	count := 0

	for num != 0 {
		num &= (num - 1)
		count++
	}

	return count
}
