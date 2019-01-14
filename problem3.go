package main

import (
	"fmt"
)

func main() {
	nbr := 600851475143
	factor := 2
	factors := []int{}

	for {
		if nbr == 1 {
			break
		}
		if nbr%factor == 0 {
			nbr /= factor
			factors = append(factors, factor)
			factor = 2
		}
		factor++
	}
	fmt.Println(factors)
}
