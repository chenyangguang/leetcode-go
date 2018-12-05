package largestime

import (
	"math/rand"
	"time"
)

/*Given an array of 4 digits, return the largest 24 hour time that can be made.

The smallest 24 hour time is 00:00, and the largest is 23:59.  Starting from 00:00, a time is larger if more time has elapsed since midnight.

Return the answer as a string of length 5.  If no valid time can be made, return an empty string.

examplame 1:
Input: [1,2,3,4]
Output: "23:41"

example 2:
Input: [5,5,5,5]
Output: ""
*/

// LargestTimeByGivenDigits return the largest time for given digits
func LargestTimeByGivenDigits(digits string) string {
	// divide into two part
	if digits == "" {
		return ""
	}
	for k, v := range digits {
		print(k, string(v), "\n")
	}
	return ""
}

// RandStr create s string by given length and given characers
func RandStr(l int, c string) string {
	seeder := rand.New(rand.NewSource(time.Now().UnixNano()))
	return func(l int, c string) string {
		ns := make([]byte, l)
		length := len(c)
		for i := range ns {
			ns[i] = c[seeder.Intn(length)]
		}
		return string(ns)
	}(l, c)
}
