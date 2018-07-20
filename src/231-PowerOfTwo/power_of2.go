// Given an integer, write a function to determine if it is a power of two.

package powof2

import (
	"math"
)

// isPowerOf2 check if a integer is power of 2
func IsPowerOf2(n int) bool {
	var tmp float64

	for i := 0; i <= n/2; i++ {
		tmp = math.Pow(2, float64(i))
		if tmp == float64(n) {
			return true
		}
		if tmp > float64(n) {
			return false
		}
	}
	return false
}
