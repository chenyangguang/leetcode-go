package powof2

import (
	"testing"
)

func TestIsPowerOf2(t *testing.T) {
	nums := []int{1, 2, 4, 8, 9, 16, 128}
	for _, v := range nums {
		err := IsPowerOf2(v)
		if !err {
			t.Error(v, " is not power of 2")
		} else {
			t.Log(v, "is power of 2")
		}
	}
}

/*
请按 ENTER 或其它命令继续
--- FAIL: TestIsPowerOf2 (0.00s)
        power_of2_test.go:14: 1 is power of 2
        power_of2_test.go:14: 2 is power of 2
        power_of2_test.go:14: 4 is power of 2
        power_of2_test.go:14: 8 is power of 2
        power_of2_test.go:12: 9  is not power of 2
        power_of2_test.go:14: 16 is power of 2
        power_of2_test.go:14: 128 is power of 2
*/
