package main

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
}
