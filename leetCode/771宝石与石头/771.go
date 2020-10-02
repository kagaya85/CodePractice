package main

import (
	"fmt"
)

func main() {

	J, S := "aA", "aAAbbbb"
	fmt.Println(numJewelsInStones(J, S))

}

func numJewelsInStones(J string, S string) (ans int) {
	jmap := make(map[rune]bool)

	for _, c := range J {
		jmap[c] = true
	}

	for _, c := range S {
		if jmap[c] == true {
			ans++
		}
	}

	return ans
}
