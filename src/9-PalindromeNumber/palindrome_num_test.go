package pn

import (
	"testing"
)

func TestPalindromeNumber(t *testing.T) {
	var num = 343
	if ok := PalindromeNumber(num); !ok {
		t.Error(num, " is not a palindrome number. ")
	}
	t.Log(num)

	num = 12345321
	if ok := PalindromeNumber(num); !ok {
		t.Error(num, " is not a palindrome number. ")
	}
	t.Log(num)

	num = -121
	if ok := PalindromeNumber(num); !ok {
		t.Error(num, " is not a palindrome number. ")
	}
	t.Log(num)
}
