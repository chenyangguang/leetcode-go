package main

import (
	"fmt"
)

func main() {
	var testIntArr = []int{2, 7, 11, 17}
	fmt.Println(twoSum(testIntArr, 9))   // [0,2]
	fmt.Println(twoSum2(testIntArr, 24)) // [1,3]
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

func twoSum2(nArr []int, t int) []int {
	ret := []int{}
	lens := len(nArr)
	for i := 0; i < lens; i++ {
		for j := 0; j < lens; j++ {
			if i != j && nArr[i]+nArr[j] == t {
				ret = append(ret, i)
				ret = append(ret, j)
				return ret
			}
		}
	}
	return ret
}
