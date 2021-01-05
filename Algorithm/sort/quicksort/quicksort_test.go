package quicksort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {

	assert.Equal(t, []int{}, QuickSort([]int{}))

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, QuickSort([]int{6, 5, 4, 3, 2, 1}))

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, QuickSort([]int{5, 6, 1, 4, 2, 3}))

	assert.Equal(t, []int{0, 0, 1, 1, 2, 5}, QuickSort([]int{5, 1, 1, 2, 0, 0}))
}
