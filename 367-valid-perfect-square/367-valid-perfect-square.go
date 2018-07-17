package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i < 99; i++ {
		var verify = func() {
			randNum := r1.Intn(1000)
			err, num := isPerfectSquare(randNum)
			if !err {
				fmt.Println("rand num ", randNum, "is none's square") // rand num  64 is 8's square
			} else {
				fmt.Println("rand num ", randNum, "is ", num, "'s square") // rand num  54 is none's square
			}
		}
		verify()
	}
}

//Given a positive integer num,
//write a function which returns True if num is a perfect square else False.
//Note: Do not use any built-in library function such as sqrt.
func isPerfectSquare(num int) (bool, int) {
	if num == 1 {
		return true, 1
	} else if num < 3 {
		return false, 0
	}
	if num > 3 {
		halfPos := num / 2
		for i := 2; i <= halfPos; i++ {
			if num == i*i {
				return true, i
			}
		}
	}
	return false, 0
}

// test

/*
rand num  945 is none's square
rand num  83 is none's square
rand num  674 is none's square
rand num  273 is none's square
rand num  558 is none's square
rand num  927 is none's square
rand num  838 is none's square
rand num  680 is none's square
rand num  158 is none's square
rand num  1 is  1 's square
rand num  231 is none's square
rand num  806 is none's square
rand num  59 is none's square
rand num  632 is none's square
rand num  576 is  24 's square
rand num  685 is none's square
rand num  557 is none's square
rand num  292 is none's square
rand num  557 is none's square
rand num  799 is none's square
rand num  811 is none's square
rand num  115 is none's square
rand num  369 is none's square
rand num  957 is none's square
rand num  340 is none's square
rand num  427 is none's square
rand num  38 is none's square
rand num  739 is none's square
rand num  10 is none's square
rand num  84 is none's square
rand num  195 is none's square
rand num  935 is none's square
rand num  294 is none's square
rand num  80 is none's square
rand num  421 is none's square
rand num  909 is none's square
rand num  802 is none's square
rand num  324 is  18 's square
rand num  13 is none's square
rand num  438 is none's square
rand num  227 is none's square
rand num  208 is none's square
rand num  178 is none's square
rand num  402 is none's square
rand num  959 is none's square
rand num  579 is none's square
rand num  778 is none's square
rand num  502 is none's square
rand num  136 is none's square
rand num  154 is none's square
rand num  577 is none's square
rand num  249 is none's square
rand num  770 is none's square
rand num  571 is none's square
rand num  867 is none's square
rand num  96 is none's square
rand num  852 is none's square
rand num  31 is none's square
rand num  374 is none's square
rand num  492 is none's square
rand num  841 is  29 's square
rand num  901 is none's square
rand num  869 is none's square
rand num  670 is none's square
rand num  403 is none's square
rand num  756 is none's square
rand num  243 is none's square
rand num  996 is none's square
rand num  6 is none's square
rand num  749 is none's square
rand num  566 is none's square
rand num  402 is none's square
rand num  293 is none's square
rand num  125 is none's square
rand num  852 is none's square
rand num  352 is none's square
rand num  821 is none's square
rand num  800 is none's square
rand num  166 is none's square
rand num  770 is none's square
rand num  431 is none's square
rand num  719 is none's square
rand num  157 is none's square
rand num  65 is none's square
rand num  344 is none's square
rand num  461 is none's square
rand num  285 is none's square
rand num  446 is none's square
rand num  387 is none's square
rand num  801 is none's square
rand num  321 is none's square
rand num  317 is none's square
rand num  306 is none's square
rand num  772 is none's square
rand num  372 is none's square
rand num  954 is none's square
rand num  784 is  28 's square
rand num  875 is none's square
*/
