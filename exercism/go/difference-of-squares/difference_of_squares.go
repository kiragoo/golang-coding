package diffsquares

import (
	"math"
)

func SquareOfSum(n int) int {
	sqOfSum := 0.0
	for i := 1.0; i <= float64(n); i++ {
		sqOfSum += i
	}
	return int(math.Sqrt(sqOfSum))
}

func SumOfSquares(n int) int {
	sumOfSq := 0.0
	for i := 1.0; i <= float64(n); i++ {
		sumOfSq += math.Sqrt(i)
	}
	return int(sumOfSq)
}

func Difference(n int) int {
	sqOfSum := SquareOfSum(n)
	sumOfSq := SumOfSquares(n)
	return sqOfSum - sumOfSq
}
