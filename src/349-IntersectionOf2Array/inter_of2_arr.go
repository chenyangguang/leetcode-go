// Given two arrays, write a function to compute their intersection.
// Example: Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2].
// Note:
// Each element in the result must be unique.
// The result can be in any order.

package inof2arr

import (
	"strconv"
)

func IntersectionOf2Array(iarr1, iarr2 []int) []int {
	var ret []int
	len1 := len(iarr1)
	len2 := len(iarr2)

	var maxLen, minLen int
	if len1 > len2 {
		maxLen = len1
		minLen = len2
	} else {
		maxLen = len2
		maxLen = len1
	}

	m := make(map[string]int)
	for i := 0; i < maxLen; i++ {
		for j := 0; j < minLen; j++ {
			if iarr1[i] != iarr2[j] {
				continue
			}
			k := strconv.Itoa(i)
			m[k] = iarr1[i]
			if _, ok := m[k]; !ok {
				ret = append(ret, iarr1[i])
			}
		}
	}
	return ret
}
