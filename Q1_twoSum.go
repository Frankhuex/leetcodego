package leetcode

func twoSum(nums []int, target int) []int {
	visited := map[int]int{}
	for i, n := range nums {
		if j, exists := visited[target-n]; exists {
			return []int{j, i}
		}
		visited[n] = i
	}
	return []int{0, 0}
}
