package main

import "fmt"

const primesDesired = 10001

func main() {
	primes := []int{2}

	primesSoFar := 1
	for i := 2; true; i++ {
		isPrime := true
		for _, prime := range primes {
			if i%prime == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primesSoFar++
			primes = append(primes, i)
		}

		if primesSoFar == primesDesired {
			fmt.Println(i)
			break
		}
	}
}
