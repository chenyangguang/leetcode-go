package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	fmt.Println("hehe")
	var nums = []int{-9, 2, 3, 4, 9, 11}

	rand.Seed(time.Now().Unix())
	randN := rand.Intn(9999)
	target := rand.Intn(9999)

	// fmt.Println(search(nums, 9))

	nums = initNums(-9999, 9999, randN)
	fmt.Println("nums=:", nums, "target:", target, "search result:", search(nums, target))

}

// search is a binary search method
// Given a sorted (in ascending order) integer array nums of n elements and a target value, write a function to search target in nums. If target exists, then return its index, otherwise return -1.
func search(nums []int, target int) (r int) {
	left := 0
	right := len(nums) - 1
	for left <= right {
		var mid = left + (right-left)/2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// initNums rand a unique array[int] from -9999 to 9999
func initNums(min, max, n int) (ret []int) {
	rand.Seed(time.Now().Unix())
	for true {
		randk := rand.Intn(2)
		randNum := rand.Intn(max)
		if randk == 1 {
			randNum = -randNum
		}
		if len(ret) > n {
			break
		}
		sort.Ints(ret)
		var exists = false
		for _, v := range ret {
			if v == randNum {
				exists = true
			}
		}
		if !exists {
			ret = append(ret, randNum)
		}
	}
	return
}
