package main

func main() {

}

func isPossible(nums []int) bool {
	numCnt := map[int]int{}
	subCnt := map[int]int{}

	for _, v := range nums {
		numCnt[v]++
	}

	for _, v := range nums {
		if numCnt[v] == 0 {
			continue
		}
		if subCnt[v-1] > 0 {
			subCnt[v-1]--
			subCnt[v]++
			numCnt[v]--
		} else if numCnt[v+1] > 0 && numCnt[v+2] > 0 {
			numCnt[v]--
			numCnt[v+1]--
			numCnt[v+2]--
			subCnt[v+2]++
		} else {
			return false
		}
	}

	return true
}
