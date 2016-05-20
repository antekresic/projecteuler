//Problem 9: https://projecteuler.net/problem=9
package main

func main() {
	a, b, c, sum := 0, 0, 0, 1000

mainloop:
	for i := sum; i > 0; i-- {
		c = i
		for j := i - 1; j > 0 && sum-j-i > 0; j-- {
			b = j
			a = sum - c - b
			if a*a+b*b == c*c {
				break mainloop
			}
		}
	}

	println(a * b * c)
}
