package leetcode

func removeElement(nums []int, val int) int {
	size := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			tmp := nums[i]
			nums[i] = nums[size]
			nums[size] = tmp
			size++
		}
	}
	return size

}
