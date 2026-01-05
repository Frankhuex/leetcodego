package leetcode

func removeDuplicates(nums []int) int {
	cur := nums[0]
	size := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] > cur {
			nums[size] = nums[i]
			cur = nums[i]
			size++
		}
	}
	return size
}
