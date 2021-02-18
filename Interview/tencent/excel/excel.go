package excel

func Excel(col string) (ans int) {
	colbyte := []byte(col)
	base := 1

	for i := len(colbyte) - 1; i >= 0; i-- {
		ans += int(colbyte[i]-'A'+1) * base
		base *= 26
	}

	return
}
