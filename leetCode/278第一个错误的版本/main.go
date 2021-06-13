package main

import "sort"

func firstBadVersion(n int) int {
	return sort.Search(n, func(i int) bool {
		return isBadVersion(i)
	})
}

func isBadVersion(version int) bool {
	return true
}
