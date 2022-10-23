func divide(a int, b int) int {
	neg := false
	if a^b < 0 {
		neg = true
	}

	if a > 0 {
		a = -a
	}

	if b > 0 {
		b = -b
	}

	var res uint32
	for a <= b {
		i := uint32(1)
		tmp := b
		for a <= tmp {
			a -= tmp
			res += i
			if tmp < math.MinInt32>>1 {
				break
			}
			tmp += tmp
			i += i
		}
	}

	cutoff := uint32(1 << (32 - 1))
	if neg && res > cutoff {
		return -int(cutoff)
	}

	if !neg && res >= uint32(1<<(32-1)) {
		return int(cutoff - 1)
	}

	if neg {
		return -int(res)
	}
	return int(res)
}