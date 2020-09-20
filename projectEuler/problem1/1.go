package main

import "fmt"

func main() {
	num1, num2 := 3, 5
	limit := 1000

	sum := 0

	for multiple1, multiple2 := num1, num2; multiple1 < limit || multiple2 < limit; {
		if multiple1 < limit {
			sum += multiple1
			multiple1 += num1
		}
		if multiple2 < limit {
			if multiple2%num1 != 0 {
				sum += multiple2
			}
			multiple2 += num2
		}
	}

	fmt.Println(sum)
}
