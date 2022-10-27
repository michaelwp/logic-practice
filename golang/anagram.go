package main

import "strings"

func isAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1 = strings.ToUpper(s1)
	s2 = strings.ToUpper(s2)

	freqS1 := map[byte]int{}
	freqS2 := map[byte]int{}

	for i, s := range s1 {
		freqS1[byte(s)]++
		freqS2[byte(s2[i])]++
	}

	for k, v := range freqS1 {
		if freqS2[k] != v {
			return false
		}
	}

	return true
}
