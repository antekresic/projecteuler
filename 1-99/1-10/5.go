//Problem 5: https://projecteuler.net/problem=5
package main

func main() {
	prod := 1
	for i := 1; i <= 20; i++ {
		prod *= i
	}

	for i := 20; i > 10; i-- {
		for {
			if isDivisible(prod / i) {
				prod /= i
				continue
			}
			break
		}
	}

	println(prod)
}

func isDivisible(number int) bool {
	for i := 20; i > 10; i-- {
		if number%i == 0 {
			continue
		}
		return false
	}
	return true
}
