package main

func isSubSequence(arr []int32, sArr []int32) bool {
	if len(sArr) <= 0 || sArr == nil || len(arr) <= 0 || arr == nil {
		return false
	}

	sArrId := 0

	for _, a := range arr {
		if a == sArr[sArrId] {
			sArrId++
		}
	}

	return sArrId == len(sArr)
}
