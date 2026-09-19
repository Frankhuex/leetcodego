package leetcode

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})
	var sum int64 = 0
	for i, n := range happiness {
		if i == k {
			return sum
		}
		sum += int64(max(0, n-i))
	}
	return sum
}
