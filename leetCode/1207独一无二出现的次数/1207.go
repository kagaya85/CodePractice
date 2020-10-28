package main

func main() {

}

func uniqueOccurrences(arr []int) bool {
	count := make(map[int]int)

	for _, num := range arr {
		count[num]++
	}

	cTimes := map[int]struct{}{}

	for _, v := range count {
		if _, ok := cTimes[v]; ok {
			return false
		}
		cTimes[v] = struct{}{}
	}

	return true
}
