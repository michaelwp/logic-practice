package main

type coord struct {
	x, y int32
}

func spiralStep(matrix [][]int32) [][]int32 {
	var res = [][]int32{}

	var rLimit int32 = int32(len(matrix[0])) - 1
	var bLimit int32 = int32(len(matrix)) - 1
	var lLimit int32 = 0
	var tLimit int32 = 0

	matrixLen := (len(matrix[0]) * len(matrix))
	pos := coord{x: 0, y: 0}
	limit := &rLimit

	var elRes = []int32{}

	for i := 0; i < matrixLen; i++ {
		switch limit {
		case &bLimit:
			elRes = append(elRes, matrix[pos.y][pos.x])
			if pos.y == *limit {
				limit = &lLimit
				res = append(res, elRes)
				elRes = []int32{}
				rLimit--
				pos.x--
			} else {
				pos.y++
			}
		case &lLimit:
			elRes = append(elRes, matrix[pos.y][pos.x])
			if pos.x == *limit {
				limit = &tLimit
				res = append(res, elRes)
				elRes = []int32{}
				bLimit--
				pos.y--
			} else {
				pos.x--
			}
		case &tLimit:
			elRes = append(elRes, matrix[pos.y][pos.x])
			if pos.y == *limit {
				limit = &rLimit
				res = append(res, elRes)
				elRes = []int32{}
				lLimit++
				pos.x++
			} else {
				pos.y--
			}
		default:
			elRes = append(elRes, matrix[pos.y][pos.x])
			if pos.x == *limit {
				limit = &bLimit
				res = append(res, elRes)
				elRes = []int32{}
				tLimit++
				pos.y++
			} else {
				pos.x++
			}
		}
	}

	return res
}
