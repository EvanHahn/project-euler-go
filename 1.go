package main

import "fmt"

const max = 1000

func main() {
	sum := 0

	for i := 1; i < max; i++ {
		if ((i % 3) == 0) || ((i % 5) == 0) {
			sum += i
		}
	}

	fmt.Println(sum)
}
