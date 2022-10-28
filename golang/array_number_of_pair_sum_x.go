package main

import (
	"fmt"
	"sort"
)

func numberOfPairSumX(arr []int, sum int) (res []string) {
	res = []string{}

	// steps :
	// 1. sorting the `arr` asc
	// ------------------------
	sort.Ints(arr)

	// 2. validation :
	// - if the smallest `number` greater than `sum` then return 0
	// -----------------------------------------------------------
	if arr[0] >= sum {
		return res
	}

	// 3. iterating over the `arr` :
	// ------------------------------------------------------------------
	for _, n := range arr {

		// 	* if (sum - n) < n then break/ return the result
		if (sum - n) < n {
			break
		}

		//  * if (sum -n) == n then skip/ continue
		if (sum - n) == n {
			continue
		}

		// 	* if (sum - n) exist in arr then add the number of pair to result
		if isNumberExist(sum-n, arr) {
			res = append(res, fmt.Sprintf(`{%d,%d}`, n, sum-n))
		}
	}

	return res
}

func isNumberExist(number int, arr []int) bool {

	// 1. copy `arr` to new slice
	var newArr = make([]int, 0)
	newArr = append(newArr, arr...)

	// 2. sorting the `newArr` desc
	sort.SliceStable(newArr, func(i, j int) bool {
		return newArr[i] > newArr[j]
	})

	// 3. iterating over the `newArr`
	// if n <= number then break
	for _, n := range newArr {
		if n < number {
			break
		}

		if n == number {
			return true
		}
	}

	return false
}
