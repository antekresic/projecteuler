//Problem 4: https://projecteuler.net/problem=4
package main

func main() {
	i := 1000000
mainloop:
	for {
		i--
		if !isPalindromeNumber(i) {
			continue
		}

		for j := 999; j > 99; j-- {
			if i%j == 0 && i/j < 1000 {
				break mainloop
			}
		}
	}

	println(i)
}

func isPalindromeNumber(number int) bool {
	original := number
	reverse := 0

	for number != 0 {
		reverse = reverse*10 + number%10
		number /= 10
	}

	return original == reverse
}
