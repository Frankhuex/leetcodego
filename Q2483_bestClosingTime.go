package leetcode

import (
	"fmt"
	"math"
)

func bestClosingTime(customers string) int {
	if len(customers) == 0 {
		return 0
	}
	leftN := make([]int, len(customers)+1)
	rightY := make([]int, len(customers)+1)
	for i, c := range customers {
		if i == 0 {
			if c == 'N' {
				leftN[i] = 1
			}
		} else if c == 'N' {
			leftN[i] = leftN[i-1] + 1
		} else {
			leftN[i] = leftN[i-1]
		}
	}
	leftN[len(customers)] = leftN[len(customers)-1]

	for i := len(customers) - 1; i >= 0; i-- {
		c := customers[i]
		if c == 'Y' {
			rightY[i] = rightY[i+1] + 1
		} else {
			rightY[i] = rightY[i+1]
		}
	}

	fmt.Println(leftN)
	fmt.Println(rightY)

	minCost := math.MaxInt64
	bestI := 0
	for i := 0; i <= len(customers); i++ {
		var cost int
		if i == 0 {
			cost = rightY[i]
		} else if i == len(customers) {
			cost = leftN[i-1]
		} else {
			cost = leftN[i-1] + rightY[i]
		}
		if cost < minCost {
			bestI = i
			minCost = cost
		}
	}
	return bestI
}
