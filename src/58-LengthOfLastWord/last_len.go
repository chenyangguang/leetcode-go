package ll

import (
	"strings"
)

// Given a string s consists of upper/lower-case alphabets and
// empty space characters ' ', return the length of last word in the string.
// If the last word does not exist, return 0.
// Note: A word is defined as a character sequence consists of
// non-space characters only.
// see more from https://leetcode.com/problems/length-of-last-word/description/
func LengthOfLastWord(s string) int {
	arr := strings.Split(s, " ")

	return len(arr[len(arr)-1])
}
