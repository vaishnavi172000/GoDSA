package godsa

// Write an algorithm to determine if a number n is happy.

// A happy number is a number defined by the following process:

// Starting with any positive integer, replace the number by the sum of the squares of its digits.
// Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy.
// Return true if n is a happy number, and false if not.

// Example 1:

// Input: n = 19
// Output: true
// Explanation:
// 12 + 92 = 82
// 82 + 22 = 68
// 62 + 82 = 100
// 12 + 02 + 02 = 1
// Example 2:

// Input: n = 2
// Output: false
func isHappy(n int) bool {
	m1 := make(map[int64]struct{})
	var n1 int64
	n1 = int64(n)
	for n1 != 1 {
		n1 = sumDigitSquares(n1)
		if _, ok := m1[n1]; ok {

			return false
		}
		m1[n1] = struct{}{}

	}
	return true

}

func sumDigitSquares(n int64) int64 {
	var sum int64
	for n > 0 {
		m := n % 10
		sum = sum + m*m
		n = n / 10
	}
	return sum

}
