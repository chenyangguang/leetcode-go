package main

import (
	"fmt"
)

func main() {
	fmt.Println("smt")
	var testIntArr = []int{2, 7, 11, 17}
	fmt.Println(twoSum(testIntArr, 9)) // [0,1]
}

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// 	You may assume that each input would have exactly one solution, and you may not use the same element twice.
// example:
// Given nums = [2, 7, 11, 15], target = 9,
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

func twoSum(nums []int, target int) []int {
	var ret = []int{}

	for k, v := range nums {
		for kk, vv := range nums {
			if target == v+vv {
				ret = append(ret, k)
				ret = append(ret, kk)
				return ret
			}
		}
	}
	return ret
}
