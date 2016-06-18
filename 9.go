package main

import "fmt"

const sum = 1000

func main() {
	for a := 1; a < sum; a++ {
		for b := a + 1; b < sum; b++ {
			c := sum - a - b
			if c < b {
				break
			}

			if (a+b+c == sum) && (a*a+b*b == c*c) {
				fmt.Println(a * b * c)
				break
			}
		}
	}
}
