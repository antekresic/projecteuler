//Problem 3: https://projecteuler.net/problem=3
package main

func main() {
	result := 0
	number := 600851475143
	i := 1
	notPrime := false
	var primes []int

	for {
		notPrime = false
		i++
		for _, v := range primes {
			if i%v == 0 {
				notPrime = true
				break
			}
		}

		if notPrime {
			continue
		}

		primes = append(primes, i)

		for {
			if number%i == 0 {
				number = number / i
				result = i
			} else {
				break
			}
		}

		if number == i || number == 1 {
			break
		}
	}

	println(result)
}
