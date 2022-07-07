package main

import "strings"

func numberOfOccurrences(word string) (res map[string]int) {
	res = map[string]int{}

	// steps :
	// 1. transform word to lowercase
	word = strings.ToLower(word)

	// 2. iterate over word.
	// - convert char from `rune` to `string`
	// - Check if res[char] == 0 then add char to map with value 1 else value++
	for _, char := range word {
		charStr := string(char)
		if res[charStr] == 0 {
			res[charStr] = 1
		} else {
			res[charStr]++
		}
	}

	return res
}
