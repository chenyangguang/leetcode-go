package inof2arr

import (
	"testing"
)

func TestIntersectionOf2Array(t *testing.T) {
	arr_1 := []int{1, 1, 2, 2}
	arr_2 := []int{2, 2}

	res := IntersectionOf2Array(arr_1, arr_2)
	for _, v := range res {
		if v == 2 {
			t.Log("arr_1:", arr_1, " and arr_2:", arr_2, " have intersection partition")
			break
		}
		t.Error("arr_1:", arr_1, " and arr_2:", arr_2, " have no intersection partition")
	}
}
