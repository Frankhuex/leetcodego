package leetcode

func nextPermutation(nums []int) {
	i := len(nums) - 1
	for i > 0 && nums[i] <= nums[i-1] {
		i--
	}
	for j := i; j < (i+len(nums))/2; j++ {
		tmp := nums[j]
		nums[j] = nums[i+len(nums)-1-j]
		nums[i+len(nums)-1-j] = tmp
	}
	if i == 0 {
		return
	}
	for k := i; k < len(nums); k++ {
		if nums[k] > nums[i-1] {
			tmp := nums[i-1]
			nums[i-1] = nums[k]
			nums[k] = tmp
			break
		}
	}

}
