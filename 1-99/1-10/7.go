//Problem 7: https://projecteuler.net/problem=7
package main

func main() {
	var primes []int
	position := 10001
	i := 1

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

		if len(primes) == position {
			break
		}
	}

	println(primes[len(primes)-1])
}
