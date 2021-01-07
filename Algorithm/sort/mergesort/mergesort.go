package mergesort

func MergeSort(arr []int) []int {
	tmp := make([]int, len(arr))

	var mergeSort func([]int, int, int) []int
	mergeSort = func(arr []int, left int, right int) []int {
		if right-left <= 1 {
			return arr
		}
		mid := (right + left) >> 1
		mergeSort(arr, left, mid)
		mergeSort(arr, mid, right)

		for p, q, i := left, mid, left; i < right; i++ {
			if q >= right || (p < mid && arr[p] < arr[q]) {
				tmp[i] = arr[p]
				p++
			} else {
				tmp[i] = arr[q]
				q++
			}
		}

		for i := left; i < right; i++ {
			arr[i] = tmp[i]
		}

		return arr
	}

	return mergeSort(arr, 0, len(arr))
}
