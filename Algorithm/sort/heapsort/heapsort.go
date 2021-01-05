package heapsort

func HeapSort(arr []int) []int {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		// 从最后一个父节点向前迭代调整为大顶堆
		justify(arr, i, n)
	}

	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		justify(arr, 0, i)
	}

	return arr
}

// 将当前的根节点向下调整到合适的位置
func justify(arr []int, start int, end int) {
	parent := start
	child := parent*2 + 1

	for child < end {
		// 选择较大的子节点
		if child+1 < end && arr[child] < arr[child+1] {
			child++
		}
		if arr[parent] > arr[child] {
			// 父节点大于较大的子节点，本次调整结束
			return
		} else {
			arr[parent], arr[child] = arr[child], arr[parent]
			parent = child
			child = parent*2 + 1
		}
	}
}
