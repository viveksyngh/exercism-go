package diffsquares

//SumOfSquares returns sum of squares of n numbers
func SumOfSquares(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += (i * i)
	}
	return sum
}

//SquareOfSums returns suqare of sum of n numbers
func SquareOfSums(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

//Difference returns difference of square of sums and sum of squares
func Difference(n int) int {

	return SquareOfSums(n) - SumOfSquares(n)
}
