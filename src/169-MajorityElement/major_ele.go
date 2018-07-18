package me

import (
	"strconv"
)

// Given an array of size n,
// find the majority element.
// The majority element is the element that appears more than  n/2  times.
// You may assume that the array is non-empty and
// the majority element always exist in the array.
// see more from https://leetcode.com/problems/majority-element/description/
func MajorityElement(nums []int) int {
	lens := len(nums)
	m := make(map[string]int)
	for _, v := range nums {
		tmpk := strconv.Itoa(v)
		if _, ok := m[tmpk]; !ok {
			m[tmpk] = 1
		} else {
			m[tmpk] += 1
		}
		if m[tmpk] > lens/2 {
			return v
		}
	}
	return 0
}
