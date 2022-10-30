package main

import "strings"

func isAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1 = strings.ToUpper(s1)
	s2 = strings.ToUpper(s2)

	freqS1 := map[rune]int{}

	for _, s := range s1 {
		freqS1[s]++
	}

	for _, s := range s2 {
		if freqS1[s] <= 0 {
			return false
		}

		freqS1[s]--
	}

	return true
}

func sherlockAndAnagrams(s string) int32 {
	if s == "" {
		return 0
	}

	var total int32 = 0

	for n := 0; n < len(s)-1; n++ {
		for i := 0; i < (len(s)-1)-n; i++ {
			for j := i + 1; j < (len(s))-n; j++ {
				strI := s[i : (i+1)+n]
				strJ := s[j : (j+1)+n]

				if isAnagram(strI, strJ) {
					total++
				}
			}
		}
	}

	return total
}
