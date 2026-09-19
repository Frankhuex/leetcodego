package leetcode

import "math"

func largestPrime(n int) int {
	if n < 2 {
		return 0
	}
	isPrime := func(n int) bool {
		for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
	sum := 2
	maxPrime := 2
	for i := 3; i < n; i++ {
		if isPrime(i) {
			sum += i
			if sum > n {
				return maxPrime
			}
			if isPrime(sum) {
				maxPrime = sum
			}
		}
	}
	return maxPrime
}
