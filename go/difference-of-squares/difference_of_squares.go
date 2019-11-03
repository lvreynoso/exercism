// difference_of_squares.go

package diffsquares 

// SquareOfSum - given an int N, SquareOfSum will sum the first N integers,
// square the sum, and return the result.
func SquareOfSum(n int) (int) {
    sum := 0

    for integer := 1; integer <= n; integer++ {
        sum += integer
    }

    return sum * sum
}

// SumOfSquares - given an int N, SumOfSquares will sum together the squares of the
// first N integers.
func SumOfSquares(n int) (int) {
    sum := 0

    for integer := 1; integer <= n; integer++ {
        sum += integer * integer
    }

    return sum
}

// Difference - given an int N, Difference will return the difference between
// the square of the sum of the first N integers and the sum of the
// squares of the first N integers.
func Difference(n int) (int) {
    return SquareOfSum(n) - SumOfSquares(n)
}