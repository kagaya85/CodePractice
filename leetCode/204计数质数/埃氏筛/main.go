package main

func countPrimes(n int) (ans int) {
	isPrime := make([]bool, n)

	for i := range isPrime {
		isPrime[i] = true
	}

	for i := 2; i < n; i++ {
		if isPrime[i] {
			ans++
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return
}
