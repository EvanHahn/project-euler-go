package main

import "fmt"

const max = 100

func sumOfSquares() int {
	sum := 0
	for i := 1; i <= max; i++ {
		sum += i * i
	}
	return sum
}

func squareOfSum() int {
	sum := 0
	for i := 1; i <= max; i++ {
		sum += i
	}
	return sum * sum
}

func main() {
	fmt.Println(squareOfSum() - sumOfSquares())
}
