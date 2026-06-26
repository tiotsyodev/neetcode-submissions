func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1

	for r > l {
		mid := l + (r - l) / 2

		if nums[mid] >= nums[r] {
			l = mid + 1
		} else if nums[mid] < nums[l] {
			r = mid
		} else {
			return nums[l]
		}
	}

	return nums[l]
}
