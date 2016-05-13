//Problem 3: https://projecteuler.net/problem=3
package main

func main() {
	result := 0
	number := 600851475143
	i := 1
	var primes []int

mainloop:
	for {
		i++
		//check if i is a prime, continue if not
		for _, v := range primes {
			if i%v == 0 {
				continue mainloop
			}
		}

		primes = append(primes, i)

		for {
			if number%i != 0 {
				break
			}

			number = number / i
			result = i
		}

		if number == 1 {
			break
		}
	}

	println(result)
}
