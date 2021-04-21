package diffsquares

import (
	"math"
)

func SquareOfSum(n int) int{
	sqOfSum := 0.0
	for i:=1.0;i<=float64(n);i++ {
		sqOfSum += i
	}
	t := math.Pow(sqOfSum,2)
	return int(t)

}

func SumOfSquares(n int) int{
	sumOfSq := 0.0
	for i:=1.0;i<=float64(n);i++ {
		sumOfSq += math.Pow(i,2)
	}
	t := sumOfSq
	return int(t)
}

func Difference(n int) int{
	sqOfSum := SquareOfSum(n)
	sumOfSq := SumOfSquares(n)
	return sqOfSum-sumOfSq
}