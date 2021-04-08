package array

func FindMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	for {
		if left >= right {
			break
		}
		if nums[left] <= nums[right] {
			return nums[left]
		}
		mid := (left + right) / 2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid] < nums[mid-1] {
			return nums[mid]
		}
		if nums[mid] > nums[0] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return nums[left]
}
