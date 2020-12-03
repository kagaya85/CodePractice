package main

func main() {

}

func countPrimes(n int) (ans int) {
	for i := 2; i < n; i++ {
		if isPrime(i) {
			ans++
		}
	}
	return
}

func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
