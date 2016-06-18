package main

import "fmt"

const min = 1
const max = 20

func isDivisibleByAll(n int) bool {
	for i := min; i <= max; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true
}

func main() {
	for i := max; true; i += max {
		if isDivisibleByAll(i) {
			fmt.Println(i)
			break
		}
	}
}
