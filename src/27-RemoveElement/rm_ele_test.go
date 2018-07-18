package re

import (
	"testing"
)

func TestRemoveEle(t *testing.T) {
	var nums = []int{3, 2, 2, 3}
	len := RemoveEle(nums, 3)
	if len != 2 {
		t.Error("fail")
	}
	t.Log(len)

	var nums_2 = []int{3, 2, 2, 3, 0, 1, 2}
	len_2 := RemoveEle(nums_2, 2)
	if len_2 != 4 {
		t.Error("fail")
	}
	t.Log(len_2)
}
