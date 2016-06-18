package main

import (
	"fmt"
	"strconv"
)

const min = 100
const max = 999

func isPalindrome(n int) bool {
	str := strconv.Itoa(n)
	length := len(str)

	for i := 0; i < (length / 2); i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}

	return true
}

func main() {
	largest := 0

	for i := max; i >= min; i-- {
		for j := max; j >= min; j-- {
			product := i * j
			if isPalindrome(i*j) && (largest < product) {
				largest = product
			}
		}
	}

	fmt.Println(largest)
}
