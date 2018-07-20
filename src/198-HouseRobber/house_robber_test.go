package rob

import (
	"testing"
)

func TestRob(t *testing.T) {
	var amount = []int{1, 2, 3, 1}
	max := Rob(amount)
	if max != 4 {
		t.Error("rob max fail")
	}
	t.Log(max)
}
