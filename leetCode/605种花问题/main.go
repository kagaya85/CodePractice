package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := range flowerbed {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			n--
			if n <= 0 {
				return true
			}
			flowerbed[i] = 1
		}
	}
	return n <= 0
}
