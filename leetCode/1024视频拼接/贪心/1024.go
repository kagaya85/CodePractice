package main

func main() {

}

func videoStitching(clips [][]int, T int) (ans int) {
	maxRight := make([]int, T)
	for _, clip := range clips {
		left, right := clip[0], clip[1]
		if left < T && right > maxRight[left] {
			maxRight[left] = right
		}
	}

	pre, next := 0, 0
	for i, max := range maxRight {
		if i > next {
			return -1
		}
		if max > next {
			next = max
		}
		if i == pre {
			pre = next
			ans++
		}
	}

	return
}
