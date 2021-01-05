package quicksort

func QuickSort(arr []int) []int {
	return quickSort(arr, 0, len(arr))
}

func quickSort(arr []int, left int, right int) []int {
	if left < right-1 {
		pivot := partition(arr, left, right)
		quickSort(arr, left, pivot)
		quickSort(arr, pivot+1, right)
	}
	return arr
}

func partition(arr []int, left int, right int) int {
	pivot := left
	index := left+1

	for i := index; i < right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}

	arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
	return index-1
}