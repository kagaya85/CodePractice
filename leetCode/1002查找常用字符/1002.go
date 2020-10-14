package main

import "math"

func main() {

}

func commonChars(A []string) (ans []string) {
	charMinFreq := make([]int, 26)
	freq := make([]int, 26)

	for i := range charMinFreq {
		charMinFreq[i] = math.MaxInt32
	}

	for _, w := range A {
		for _, c := range w {
			freq[c-'a']++
		}
		for i, n := range freq {
			if n < charMinFreq[i] {
				charMinFreq[i] = n
			}
			freq[i] = 0
		}
	}

	for c := byte(0); c < 26; c++ {
		for i := 0; i < charMinFreq[c]; i++ {
			ans = append(ans, string(c+'a'))
		}
	}

	return
}
