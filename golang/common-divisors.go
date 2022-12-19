package main

import "sort"

func CommonDivisors(x ...int) []int {
	if len(x) < 2 {
		return nil
	}

	var res = make([]int, 0)
	sort.Ints(x)

	if x[0] == 0 {
		return nil
	}

	var setMap = map[int]bool{}
	var pNumbers = primeNumbers(x[0])

	// optional action
	defer func() {
		setMap = nil
		pNumbers = nil
	}()

	for i := 1; i < len(x); i++ {
		for _, p := range pNumbers {
			if x[i]%p == 0 && !setMap[p] {
				setMap[p] = true
				res = append(res, p)
			}
		}
	}

	return res
}

// find the prime number for the minimum number
func primeNumbers[T int | int32 | int64](x T) []T {
	var res = make([]T, 0)
	var i T = 2
	var div = x / i

	res = append(res, 1)

	for i <= div {
		if x%i == 0 {
			res = append(res, i)
		}

		i++
	}

	res = append(res, x)
	return res
}
