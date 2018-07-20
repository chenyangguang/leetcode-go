// Given a pattern and a string str, find if str follows the same pattern.
// Here follow means a full match,
// such that there is a bijection between a letter in pattern and a non-empty word in str
// Input: pattern = "abba", str = "dog cat cat dog"
// Output: true
// Input:pattern = "abba", str = "dog cat cat fish"
// Output: false
// Input: pattern = "abba", str = "dog dog dog dog"
// Output: false
// note:
// assume pattern contains only lowercase letters,
// and str contains lowercase letters separated by a single space.
// see more detail from https://leetcode.com/problems/word-pattern/description/

package wordpat

import (
	"strings"
)

func WordPattern(pattern, str string) bool {
	pat := strings.Split(pattern, "")
	strs := strings.Split(strings.TrimSpace(str), " ")

	length := len(pat)
	if length != len(strs) {
		return false
	}

	m := make(map[string]string)
	for i := 0; i < length; i++ {
		if _, ok := m[pat[i]]; !ok {
			m[pat[i]] = strs[i]
		}
		if m[pat[i]] != strs[i] {
			return false
		}
	}

	return true
}
