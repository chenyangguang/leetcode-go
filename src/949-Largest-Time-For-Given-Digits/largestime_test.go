package largestime

import (
	"testing"
)

func Test_LargestTimeByGivenDigits(t *testing.T) {
	strs := RandStr(4, "123456789")

	print(strs, "\n")
	//t.Log(strs)
	//print(largestime)
	largesTime := LargestTimeByGivenDigits(strs)

	print(largesTime)
	//t.Log(strs, largesTime)
}
