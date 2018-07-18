package sip

import (
	"testing"
)

func TestSearchInsertPosition(t *testing.T) {
	var nums = []int{2, 3, 5, 6}
	pos := searchInsertPosition(nums, 9)
	if pos != 4 {
		t.Error("pos is wrong  postion  inserted", pos)
	}
	t.Log(pos)

	pos2 := searchInsertPosition(nums, 5)
	if pos2 != 2 {
		t.Error("pos is wrong  postion  inserted", pos2)
	}
	t.Log(pos)

	posx := searchInsertPosition(nums, 0)
	if posx != 0 {
		t.Error("pos is wrong  postion  inserted", posx)
	}
	t.Log(pos)
}
