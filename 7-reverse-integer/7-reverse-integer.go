package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	x := 1234
	y := -1234
	z := 1200
	m := -1200
	fmt.Println(x, y, z, m) //1234 -1234 1200 -1200
	fmt.Println(
		reverse2(x),
		reverse2(y),
		reverse2(z),
		reverse2(m),
	) //4321 -4321 21 -21
}

// reverse digits of an Given a 32-bit signed integer.
// Input: 123
// Output: 321
// Input: -123
// Output: -321
// Input: 120
// Output: 21
// Note:
// Assume we are dealing with an environment which could only store integers
// within the 32-bit signed integer range: [−2^31,  2^31 − 1].
// For the purpose of this problem, assume that your function returns 0
// when the reversed integer overflows.
func reverse2(x int) (r int) {
	if x == 0 {
		return x
	}
	var isNagative bool
	if x < 0 {
		isNagative = true
		x = -x
	}
	tmp := strconv.Itoa(x)
	strArr := strings.Split(tmp, "")

	lens := len(strArr)
	last := lens - 1
	for i, j := 0, last; i < j; i, j = i+1, j-1 {
		strArr[i], strArr[j] = strArr[j], strArr[i]
	}
	rstr := strings.Join(strArr[:], "")
	rInt, _ := strconv.Atoi(rstr)

	if isNagative {
		r = -rInt
	} else {
		r = rInt
	}

	return
}
