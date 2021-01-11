package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestStringWithSwaps(t *testing.T) {

	assert.Equal(t, "abcd", smallestStringWithSwaps("dabc", [][]int{[]int{0, 3}, []int{1, 2}, []int{0, 2}}))
}
