package pn

import (
	"strconv"
	"strings"
)

// PalindromeNumber Determine whether an integer is a palindrome.
// An integer is a palindrome when it reads the same backward as forward.

func PalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}
	arr := strings.Split(strconv.Itoa(x), "")
	lens := len(arr)
	for k, v := range arr {
		if k <= (lens/2) && v != arr[lens-1-k] {
			return false
		}
	}
	return true
}
