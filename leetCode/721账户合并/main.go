package main

import "sort"

func accountsMerge(accounts [][]string) (ans [][]string) {
	emailToIndex := map[string]int{}
	emailToName := map[string]string{}

	for _, account := range accounts {
		name := account[0]
		for _, email := range account[1:] {
			if _, has := emailToIndex[email]; !has {
				emailToIndex[email] = len(emailToIndex)
				emailToName[email] = name
			}
		}
	}

	ancestor := make([]int, len(emailToIndex))
	for i := range ancestor {
		ancestor[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if ancestor[x] != x {
			ancestor[x] = find(ancestor[x])
		}
		return ancestor[x]
	}

	union := func(from, to int) {
		ancestor[find(from)] = find(to)
	}

	for _, account := range accounts {
		firstIndex := emailToIndex[account[1]]
		for _, email := range account[2:] {
			union(emailToIndex[email], firstIndex)
		}
	}

	indexToEmails := map[int][]string{}
	for email, index := range emailToIndex {
		index := find(index)
		indexToEmails[index] = append(indexToEmails[index], email)
	}

	for _, emails := range indexToEmails {
		sort.Strings(emails)
		account := append([]string{emailToName[emails[0]]}, emails...)
		ans = append(ans, account)
	}
	return
}
