package main

func theNearestBlock(blocks []map[string]bool, reqs []string) int32 {
	var index int32 = 0
	var minDistance = len(blocks)
	var reqsMap = map[string]bool{}

	for _, r := range reqs {
		reqsMap[r] = true
	}

	for i, currBlock := range blocks {
		var checklist = map[string]int{}
		var subMaxDistance = -1

		leftBlock := blocks[:i]
		rightBlock := blocks[i+1:]

		for k, v := range currBlock {
			if reqsMap[k] == v {
				checklist[k] = -1
			}
		}

		limit := 0
		for d, b := range rightBlock {
			if limit == len(reqs) {
				break
			}

			for k, v := range b {
				if reqsMap[k] == v {
					if checklist[k] == 0 {
						checklist[k] = d + 1
						if d+1 > subMaxDistance {
							subMaxDistance = d + 1
						}

						limit++
					} else if (d + 1) < checklist[k] {
						checklist[k] = d + 1
						if d+1 < subMaxDistance {
							subMaxDistance = d + 1
						}

						limit++
					}
				}
			}
		}

		limit = 0
		for d := i - 1; d >= 0; d-- {
			if limit == len(reqs) {
				break
			}

			for k, v := range leftBlock[d] {
				if reqsMap[k] == v {
					if checklist[k] == 0 {
						checklist[k] = i + d
						if i-d > subMaxDistance {
							subMaxDistance = i - d
						}

						limit++
					} else if (i - d) < checklist[k] {
						checklist[k] = i - d
						if i-d < subMaxDistance {
							subMaxDistance = i - d
						}

						limit++
					}
				}
			}
		}

		if subMaxDistance < minDistance {
			minDistance = subMaxDistance
			index = int32(i)
		}
	}

	return index
}
