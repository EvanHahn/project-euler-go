package main

import (
	"fmt"
	"math"
)

const number = 600851475143

func intSqrt(n int) int {
	return int(math.Sqrt(float64(n)))
}

func isPrime(n int) bool {
	sqrt := intSqrt(n)

	for i := 2; i <= sqrt; i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

func main() {
	sqrt := intSqrt(number)

	for i := sqrt; i >= 2; i-- {
		if ((number % i) == 0) && (isPrime(i)) {
			fmt.Println(i)
			break
		}
	}
}
