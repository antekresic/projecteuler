//Problem 6: https://projecteuler.net/problem=6
package main

func main() {
	sumOfSquares := 0
	sum := 0
	for i := 1; i <= 100; i++ {
		sumOfSquares += i * i
		sum += i
	}

	println(sum*sum - sumOfSquares)
}
