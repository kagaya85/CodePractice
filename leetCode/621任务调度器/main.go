package main

func main() {

}

func leastInterval(tasks []byte, n int) int {
	taskCnt := map[byte]int{}
	for _, t := range tasks {
		taskCnt[t]++
	}

	maxExec, maxExecCnt := 0, 0
	for _, c := range taskCnt {
		if c > maxExec {
			maxExec, maxExecCnt = c, 1
		} else if c == maxExec {
			maxExecCnt++
		}
	}

	if time := (maxExec-1)*(n+1) + maxExecCnt; time > len(tasks) {
		return time
	}
	return len(tasks)
}
