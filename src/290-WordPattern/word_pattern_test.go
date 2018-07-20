package wordpat

import (
	"testing"
)

func TestWordPattern(t *testing.T) {
	pattern := "aabb"
	strs := "dog dog cat cat"
	bres := WordPattern(pattern, strs)
	if !bres {
		t.Error(pattern, " not match ", strs)
	}
	t.Log(bres)

	pattern = "aabb"
	strs = "dog dog cat fish"
	bres = WordPattern(pattern, strs)
	if !bres {
		t.Error(pattern, " not match ", strs)
	}
	t.Log(bres)

	pattern = "aaaa"
	strs = "dog dog cat dog"
	bres = WordPattern(pattern, strs)
	if !bres {
		t.Error(pattern, " not match ", strs)
	}
	t.Log(bres)

	pattern = "aaaa"
	strs = "dog dog dog dog"
	bres = WordPattern(pattern, strs)
	if !bres {
		t.Error(pattern, " not match ", strs)
	}
	t.Log(bres)
}
