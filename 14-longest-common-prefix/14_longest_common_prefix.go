package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var strs = []string{"flower", "flow", "floght", "floggg"}
	fmt.Println("The longest common prefix str is " + longestCommonPrefix(strs))
}

// longestCommonPrefix Write a function to
// find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".
func longestCommonPrefix(strs []string) string {
	var strMap = make(map[string][]string)
	for k, v := range strs {
		vArr := strings.Split(v, "")
		strMap[strconv.Itoa(k)] = vArr
	}
	var commStr string
	posi := getCommPos(strMap)
	if posi != 0 {
		for k, str := range strMap["0"] {
			if k < posi {
				commStr += str
			}
		}
	}
	return commStr
}

// get commn pos
func getCommPos(strMap map[string][]string) int {
	var retPos int
	for sk, strv := range strMap {
		var tmpStr, checkStr string
		var pos int
		isComm := true
		for k, v := range strv {
			tmpStr = strMap["0"][k]
			checkStr = v
			fmt.Println(tmpStr, checkStr)
			if tmpStr != strMap[sk][k] {
				isComm = false
				break
			}
			pos++
		}

		// stop at the not match position
		if !isComm {
			fmt.Println("break index", pos)
			return pos
		}
	}
	return retPos
}
