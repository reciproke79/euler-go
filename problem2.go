package main

import (
	"fmt"
)

func main() {
	max := 4000000
	sum := 0
	term1 := 1
	term2 := 2

	for {
		term1, term2 = nextFib(term1, term2)
		if term1%2 == 0 {
			sum += term1
		}
		if term2 > max {
			break
		}
	}
	fmt.Println(sum)
}

func nextFib(term1, term2 int) (int, int) {
	return term2, term1 + term2
}
