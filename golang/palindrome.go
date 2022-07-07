package main

import "strings"

func isPalindrome(s string) bool {
	var sOri = ""
	var sRev = ""

	// transform string to lowercase
	s = strings.ToLower(s)

	// iterate over string to get alphabet (a-z) only from string
	// then add it into 2 new string
	for _, c := range s {
		if (c >= 97 && c <= 122) || (c >= 48 && c <= 57) {
			sOri += string(c)
			sRev = string(c) + sRev
		}
	}

	// returning comparable result
	return sOri == sRev
}
