//Problem 2: https://projecteuler.net/problem=2
package main

func main() {
	sum := 2
	fib1 := 1
	fib2 := 2

	for {
		fib1, fib2 = fib2, fib1+fib2

		if fib2%2 == 0 {
			sum += fib2
		}

		if fib2 > 4000000 {
			break
		}
	}

	println(sum)
}
