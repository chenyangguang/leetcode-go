package me

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	var nums = []int{3, 2, 3}
	ele := MajorityElement(nums)
	if ele != 3 {
		t.Error("fail")
	}
	t.Log(ele)

	nums_2 := []int{2, 2, 1, 1, 2, 2, 22}
	if ele_2 := MajorityElement(nums_2); ele_2 != 2 {
		t.Error("fail")
		t.Log(ele_2)
	}
}
