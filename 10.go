package main

import "fmt"

const max = 2000000

func main() {
	primes := []int{2}

	for i := 3; i < max; i += 2 {
		isPrime := true
		for _, prime := range primes {
			if i%prime == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, i)
		}
	}

	sum := 0
	for _, prime := range primes {
		sum += prime
	}
	fmt.Println(sum)
}
