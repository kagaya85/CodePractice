package kmp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	assert.Equal(t, 1, Search("abcde", "b"))

	assert.Equal(t, 3, Search("abcde", "de"))

	assert.Equal(t, -1, Search("abcde", "xyz"))
}
