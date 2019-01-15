package main

import (
	"fmt"
)

func main() {
	n := 20
	fmt.Println(lcm(n))
}

func gcd(a, b int) int {
	if (a % b) != 0 {
		return gcd(b, a%b)
	}
	return b
}

func lcm(n int) int {
	nbr := 1
	for i := 1; i <= n; i++ {
		nbr = (nbr * i) / (gcd(nbr, i))
	}
	return nbr
}
