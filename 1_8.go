package main

import (
	"fmt"
	"strings"
)

func main() {
	type TestIsRotation struct {
		givenS1        string
		givenS2        string
		expectedResult bool
	}

	cases := []TestIsRotation{
		{
			"",
			"",
			true,
		},
		{
			"abc",
			"abcd",
			false,
		},
		{
			"abce",
			"ecbd",
			false,
		},
		{
			"abcdefg",
			"bcdefga",
			true,
		},
	}

	for i, c := range cases {
		actualResult := isRotation(c.givenS1, c.givenS2)
		if actualResult != c.expectedResult {
			fmt.Printf("test case[%d] failed!", i)
		}
	}
}

//Assume you have a method isSubstring which checks if one word is a substring of another.
// Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using only one call to isSubstring
// (i.e., “waterbottle” is a rotation of “erbottlewat”).
func isRotation(s1 string, s2 string) bool {
	//check that s1 and s2 are equal length
	if len(s1) == len(s2) {
		combined := s1 + s1
		return isSubstring(combined, s2)
	}
	return false
}

func isSubstring(s1 string, s2 string) bool {
	return strings.Contains(s1, s2)
}
