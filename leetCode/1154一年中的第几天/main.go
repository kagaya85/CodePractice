func dayOfYear(date string) (res int) {
	dateArr := strings.Split(date, "-")

	year, _ := strconv.Atoi(dateArr[0])
	flag := (year%4 == 0 && year%100 != 0) || (year%400 == 0)

	month, _ := strconv.Atoi(dateArr[1])
	day, _ := strconv.Atoi(dateArr[2])
	for i := 1; i < month; i++ {
		switch i {
		case 1, 3, 5, 7, 8, 10, 12:
			res += 31
		case 4, 6, 9, 11:
			res += 30
		default:
			if flag {
				res += 29
			} else {
				res += 28
			}
		}
	}

	return res + day
}