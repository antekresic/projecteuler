//Problem 10: https://projecteuler.net/problem=10
package main

// return the sum of primes from 2 to limit (exclusive)
// uses Sieve of Eratosthenes algorithm
func sumOfPrimes(limit int) int {
	// creates a slice of false with length of limit
	bools := make([]bool, limit)
	// implies up to the sqrt of limit
	for k := 2; k*k <= limit; k++ {
		if bools[k] != true {
			for ix := k * k; ix < limit; ix += k {
				bools[ix] = true
			}
		}
	}
	// index of remaining False in bools = a prime number
	sum := 0
	for ix := 2; ix < limit; ix++ {
		if bools[ix] != true {
			sum += ix
		}
	}
	return sum
}
func main() {
	println(sumOfPrimes(2000000))
}
