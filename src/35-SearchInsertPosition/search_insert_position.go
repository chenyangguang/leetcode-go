package sip

import (
	"sort"
)

// searchInsertPosition Given a sorted array and a target value,
// return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.
func searchInsertPosition(nums []int, target int) int {
	nums = append(nums, target)
	sort.Ints(nums)
	for k, v := range nums {
		if v == target {
			return k
		}
	}
	return 0
}
