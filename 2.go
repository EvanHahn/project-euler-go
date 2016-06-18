package main

import "fmt"

const max = 4000000

func main() {
	sum := 0

	a := 1
	b := 2

	for a < max {
		if (a % 2) == 0 {
			sum += a
		}

		aOld := a
		a = b
		b = aOld + b
	}

	fmt.Println(sum)
}
