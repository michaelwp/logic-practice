package main

import "fmt"

func binaryGap(N int) int {
	var res = 0

	// convert int to binary
	var binaries = "0"
	binaries = fmt.Sprintf("%b", N)

	// iterate over binary,
	// - if binary == "0" then max++
	// - if binary == "1" then
	// 	* if max > res then res = max
	//	* max = 0
	max := 0
	for _, b := range binaries {
		if string(b) == "0" {
			max++
		} else {
			if max > res {
				res = max
			}
			max = 0
		}
	}

	return res
}
