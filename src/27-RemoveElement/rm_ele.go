package re

import (
	"fmt"
)

// RemoveEle Given an array nums and a value val, remove all instances
// of that value in-place and return the new length.
// Do not allocate extra space for another array,
// you must do this by modifying the input array in-place with O(1) extra memory.
// The order of elements can be changed.
// It doesn't matter what you leave beyond the new length.
// see more from https://leetcode.com/problems/remove-element/description/
func RemoveEle(nums []int, t int) int {
	var newNums []int
	for _, v := range nums {
		if v == t {
			continue
		}
		newNums = append(newNums, v)
	}
	fmt.Println(newNums)
	return len(newNums)
}
