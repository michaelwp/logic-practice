package main

import "math"

func rotate(aList []interface{}, rNumber int) (result []interface{}) {

	// get length of `arr`
	arrayLength := len(aList)

	// if rNumber < 0 set isReverse = true
	// set rNumber to absolute of rotate
	isReverse := rNumber < 0
	rNumber = int(math.Abs(float64(rNumber)))

	// get division remainder (mod) rotate % arrLength
	mod := rNumber % arrayLength

	// get element start/end point
	point := arrayLength - mod

	// if reverse is true then reverse the aList
	if isReverse {
		aList = reverse(aList)
	}


	// slice aList into 2 array :
	// first array contains element array [point:]
	// second array contains element array [:point]
	array1 := aList[point:]
	array2 := aList[:point]

	// join array1 & array2 as result
	result = append(array1, array2...)

	// if isReverse is true then reverse the result
	if isReverse {
		result = reverse(result)
	}

	return result
}

func reverse(arr []interface{}) (arrRev []interface{}) {

	// get length of `arr`
	arrLength := float64(len(arr))

	// get half-length of `arr` round to ceil
	halfArrLength := math.Ceil(arrLength / float64(2))

	// iterate over `arr` to number of `halfArrLength`
	for i, j := 0, int(arrLength)-1; i < int(halfArrLength); i, j = i+1, j-1 {

		// swap the element
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}
