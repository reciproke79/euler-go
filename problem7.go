package main

import (
	"fmt"
)

func main() {
	count := 0
	currentPrime := 0

	for i := 1; count != 10001; i++ {
		if isPrime(i) {
			count++
			currentPrime = i
		}
	}
	fmt.Println(currentPrime)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
