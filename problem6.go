package main

import (
	"fmt"
	"math"
)

func main() {
	sumOfSquares := sumOfSquares(100)
	squareOfSum := squareOfSum(100)
	fmt.Println(int(squareOfSum - sumOfSquares))
}

func sumOfSquares(n int) float64 {
	sum := float64(0)
	for i := float64(1); i <= float64(n); i++ {
		sum += math.Pow(i, 2)
	}
	return sum
}

func squareOfSum(n int) float64 {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return math.Pow(float64(sum), 2)
}
