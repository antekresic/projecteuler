//Problem 1: https://projecteuler.net/problem=1
package main

func main() {
	sum := 0

	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			sum += i
			continue
		}

		if i%5 == 0 {
			sum += i
		}
	}

	println(sum)
}
