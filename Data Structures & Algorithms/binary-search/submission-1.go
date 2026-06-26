func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for r >= l{
		pivot := l + (r - l) / 2

		if nums[pivot] > target {
			r = pivot - 1
		} else if nums[pivot] < target {
			l = pivot + 1
		} else {
			return pivot
		}
	}

	return -1
}
