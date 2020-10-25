package main

func main() {

}

func longestMountain(a []int) (ans int) {
	n := len(a)
	left := 0
	for left+2 < n {
		right := left + 1
		if a[left] < a[left+1] {
			for right+1 < n && a[right] < a[right+1] {
				right++
			}
			if right < n-1 && a[right] > a[right+1] {
				for right+1 < n && a[right] > a[right+1] {
					right++
				}
				if right-left+1 > ans {
					ans = right - left + 1
				}
			} else {
				right++
			}
		}
		left = right
	}
	return
}
