package ll

import (
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	var str = "test test  mm88"
	lastLen := LengthOfLastWord(str)
	if lastLen != 4 {
		t.Error("fail")
	}
	t.Log(lastLen)

	var spastr = "hello world       "
	spaceLen := LengthOfLastWord(spastr)
	if spaceLen != 0 {
		t.Error("fail", spaceLen)
	}
	t.Log(spaceLen)
}
