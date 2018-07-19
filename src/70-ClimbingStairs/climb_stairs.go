// You are climbing a stair case. It takes n steps to reach to the top.
// Each time you can either climb 1 or 2 steps.
// In how many distinct ways can you climb to the top?
// Note: Given n will be a positive integer.
// see more from https://leetcode.com/problems/climbing-stairs/description/

package cs

import (
	"math"
)

// ClimbStairsByBruteForce calculate n steps to reach the top by brute force
func ClimbStairsByBruteForce(n int) int {
	return bruteForce(0, n)
}

func bruteForce(i, n int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	return bruteForce(i+1, n) + bruteForce(i+2, n)
}

// ClimbStairsByRecurseMemory calculate by Recursion with memorization
func ClimbStairsByRecurseMemory(n int) int {
	memo := make([]int, n+1)
	return recurseMemory(0, n, memo)
}

func recurseMemory(i, n int, memo []int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	if memo[i] > 0 {
		return memo[i]
	}
	memo[i] = recurseMemory(i+1, n, memo) + recurseMemory(i+2, n, memo)
	return memo[i]
}

// ClimbStairsByDynamicProgram calculate by dynamic programming
func ClimbStairsByDynamicProgram(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// ClimbStairsByFibnacciNumber calculate by fibnacci number
func ClimbStairsByFibnacciNumber(n int) int {
	if n == 1 {
		return 1
	}

	first := 1
	second := 2
	var third int
	for i := 3; i <= n; i++ {
		third = first + second
		first = second
		second = third
	}
	return second
}

// ClimbStairsByFibnacciFormula calculate by fibnacci formula
func ClimbStairsByFibnacciFormula(n int) int {
	var sqrt, fibn float64

	sqrt = math.Sqrt(5)
	fibn = math.Pow((float64(1)+sqrt)/2, float64(n+1)) - math.Pow((float64(1)-sqrt)/2, float64(n+1))

	return (int)(fibn / sqrt)
}
