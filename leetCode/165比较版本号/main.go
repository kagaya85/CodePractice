func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	l1, l2 := len(v1), len(v2)

	if l1 > l2 {
		for i := 0; i < l1-l2; i++ {
			v2 = append(v2, "0")
		}
	}

	if l1 < l2 {
		for i := 0; i < l2-l1; i++ {
			v1 = append(v1, "0")
		}
	}

	for i := 0; i < max(l1, l2); i++ {
		n1, _ := strconv.Atoi(v1[i])
		n2, _ := strconv.Atoi(v2[i])
		if n1 == n2 {
			continue
		} else if n1 > n2 {
			return 1
		} else {
			return -1
		}
	}

	return 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}