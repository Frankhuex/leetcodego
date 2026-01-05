package leetcode

func majorityElement(nums []int) int {
	count := 0
	candidate := nums[0]
	for _, n := range nums {
		if count == 0 {
			candidate = n
		}
		if n == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
