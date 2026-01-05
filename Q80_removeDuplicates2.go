package leetcode

func removeDuplicates2(nums []int) int {
	i, j := 0, 1
	dupl := 1
	cur := nums[0]
	for j < len(nums) {
		if nums[j] == cur {
			dupl++
			if dupl == 3 {
				nums[i] = cur
				nums[i+1] = cur
				i += 2
			}
		} else {
			if dupl < 3 {
				for k := i; k < i+dupl; k++ {
					nums[k] = cur
				}
				i += dupl
			}
			dupl = 1
			cur = nums[j]
		}
		j++
	}
	if dupl < 3 {
		for k := i; k < i+dupl; k++ {
			nums[k] = cur
		}
		i += dupl
	}
	return i
}
