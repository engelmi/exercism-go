package diffsquares

func SquareOfSum(n int) int {
	sqOfSum := 0
	for i := 1; i < n+1; i++ {
		sqOfSum += i
	}
	return sqOfSum * sqOfSum
}

func SumOfSquares(n int) int {
	sumOfSq := 0
	for i := 1; i < n+1; i++ {
		sumOfSq += i * i
	}
	return sumOfSq
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
