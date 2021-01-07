package mergesort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapSort(t *testing.T) {

	assert.Equal(t, []int{}, MergeSort([]int{}))

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, MergeSort([]int{6, 5, 4, 3, 2, 1}))

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, MergeSort([]int{5, 6, 1, 4, 2, 3}))

	assert.Equal(t, []int{0, 0, 1, 1, 2, 5}, MergeSort([]int{5, 1, 1, 2, 0, 0}))
}
