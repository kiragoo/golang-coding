package lc

func Search(nums []int, target int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}
	length := len(nums)
	var start = 0
	end := length - 1
	for {
		if start > end {
			break
		}
		mid := start + (end-start)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[start] {
			start++
			continue
		}
		if nums[start] <= nums[mid] {
			if target >= nums[start] && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if target <= nums[end] && target > nums[mid] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return false
}
