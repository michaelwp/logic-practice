package main

import (
	"math"
	"sort"
)

/*
Find the minimum difference between group 0 and group 1, guaranteed 0, 1 will occur in array `a`
Example :
 	a = [0, 1, 1, 0, 0, 1]
	b = [10, 2, 3, 1, 7, 18]

	group0 = [10, 1, 7]
	group1 = [2, 3, 18]

	answer = 1

Explain :
	10 - 2 = 8
	2 - 1 = 1
	7 - 2 = 5
	10 - 3 = 7
	3 - 1 = 2
	7 - 3 = 4
	18 - 10 = 8
	18 - 1 = 17
	18 - 7 = 11

	then answer is 1, because 1 is the minimum difference between group0 and group1

requirement :
	the result must be positive integer

constraint :
	duration must less than or equal to 4 second

specs :
	0 >= b(n) <= 1000000
	len(a) == len(b)
	2 >= len(b) <= 1000
*/

func MinDifference(a, b []int) int {

	// grouping
	var group0 = make([]int, 0)
	var group1 = make([]int, 0)

	for index, value := range a {
		if value == 0 {
			group0 = append(group0, b[index])
		} else {
			group1 = append(group1, b[index])
		}
	}

	// sorting descending
	sortingDesc(group0)
	sortingDesc(group1)

	var res = int(math.Abs(float64(group0[0] - group1[0])))

	for _, vGroup0 := range group0 {
		subRes := 0

		for _, vGroup1 := range group1 {
			subRes = vGroup0 - vGroup1

			if subRes == 0 {
				return 0
			}

			if subRes > -1 {
				break
			}
		}

		positiveInt := int(math.Abs(float64(subRes)))

		if res > positiveInt {
			res = positiveInt
		}
	}

	return res
}

func sortingDesc(slice []int) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
}
