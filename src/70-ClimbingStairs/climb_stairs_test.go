package cs

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	var n = 6
	nums := ClimbStairs(n)
	if nums != 8 {
		t.Error("fail")
	}
	t.Log(nums)
}
