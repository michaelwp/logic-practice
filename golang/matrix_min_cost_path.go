package main

type pos struct {
	x, y int
}

func minCostPath(cost [][]int) int {
	currPos := &pos{
		x: len(cost[0]) - 1,
		y: len(cost) - 1,
	}

	horz := &pos{x: 0, y: 0}
	vert := &pos{x: 0, y: 0}
	diag := &pos{x: 0, y: 0}

	minCostSum := cost[currPos.y][currPos.x] + cost[0][0]

	for currPos.x > 0 && currPos.y > 0 {

		// validate the bottom limit for x and y
		stepX, stepY := currPos.x-1, currPos.y-1

		if stepX < 0 {
			stepX = 0
		}

		if stepY < 0 {
			stepY = 0
		}

		// assign the steps
		horz.x, horz.y = stepX, currPos.y
		vert.x, vert.y = currPos.x, stepY
		diag.x, diag.y = stepX, stepY

		// find the minCost
		if cost[horz.y][horz.x] > cost[vert.y][vert.x] {
			horz.x, horz.y = vert.x, vert.y
		}

		if cost[horz.y][horz.x] > cost[diag.y][diag.x] {
			horz.x, horz.y = diag.x, diag.y
		}

		// set the minimal cost summary
		minCostSum += cost[horz.y][horz.x]

		// set the current position
		currPos.x, currPos.y = horz.x, horz.y
	}

	return minCostSum
}
