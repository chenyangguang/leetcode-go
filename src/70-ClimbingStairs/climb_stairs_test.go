package cs

import (
	"testing"
)

func TestClimbStairsByBruteForce(t *testing.T) {
	// brute force
	var n = 2
	nums := ClimbStairsByBruteForce(n)
	if nums != 2 {
		t.Error("fail")
	}
	t.Log(nums)

	n = 3
	nums = ClimbStairsByBruteForce(n)
	if nums != 3 {
		t.Error("fail")
	}
	t.Log(nums)

	n = 4
	nums = ClimbStairsByBruteForce(n)
	if nums != 5 {
		t.Error("fail")
	}
	t.Log(nums)

	n = 5
	nums = ClimbStairsByBruteForce(n)
	if nums != 8 {
		t.Error("fail")
	}
	t.Log(nums)

	// recurse memory
	n = 3
	nums = ClimbStairsByRecurseMemory(n)
	if nums != 3 {
		t.Error("fail")
	}
	t.Log(nums)

	n = 4
	nums = ClimbStairsByRecurseMemory(n)
	if nums != 5 {
		t.Error("fail")
	}
	t.Log(nums)

	n = 5
	nums = ClimbStairsByRecurseMemory(n)
	if nums != 8 {
		t.Error("fail")
	}
	t.Log(nums)

	// dynamoic programming
	n = 3
	nums = ClimbStairsByDynamicProgram(n)
	if nums != 3 {
		t.Error("fail")
	}
	t.Log(nums)

	n = 4
	nums = ClimbStairsByDynamicProgram(n)
	if nums != 5 {
		t.Error("fail")
	}
	t.Log(nums)

	n = 5
	nums = ClimbStairsByDynamicProgram(n)
	if nums != 8 {
		t.Error("fail")
	}
	t.Log(nums)

	// fibnacci number
	n = 3
	nums = ClimbStairsByFibnacciNumber(n)
	if nums != 3 {
		t.Error("fail")
	}
	t.Log(nums)

	n = 4
	nums = ClimbStairsByFibnacciNumber(n)
	if nums != 5 {
		t.Error("fail")
	}
	t.Log(nums)

	n = 5
	nums = ClimbStairsByFibnacciNumber(n)
	if nums != 8 {
		t.Error("fail")
	}
	t.Log(nums)

	// fibnacci formula testing
	n = 3
	nums = ClimbStairsByFibnacciFormula(n)
	if nums != 3 {
		t.Error("fail")
	}
	t.Log(nums)

	n = 4
	nums = ClimbStairsByFibnacciFormula(n)
	if nums != 5 {
		t.Error("fail")
	}
	t.Log(nums)

	n = 5
	nums = ClimbStairsByFibnacciFormula(n)
	if nums != 8 {
		t.Error("fail")
	}
	t.Log(nums)
}
