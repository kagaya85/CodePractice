package main

func main() {

}

func longestMountain(A []int) (ans int) {
	flag := 0
	tmpLen := 1

	for i, h := range A {
		if i+1 >= len(A) {
			if tmpLen > ans && flag == 1 {
				ans = tmpLen
			}
			break
		}
		next := A[i+1]
		if flag == 0 {
			if h < next {
				tmpLen++
			} else if h > next && tmpLen > 1 {
				flag = 1
				tmpLen++
			} else {
				tmpLen = 1
			}
		} else {
			if h > next {
				tmpLen++
			} else if h <= next {
				if tmpLen > ans {
					ans = tmpLen
				}
				flag = 0
				if h != next {
					tmpLen = 2
				} else {
					tmpLen = 1
				}
			}
		}
	}

	if ans < 3 {
		return 0
	}
	return
}
