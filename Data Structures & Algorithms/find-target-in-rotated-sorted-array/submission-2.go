func search(nums []int, target int) int {
	l, r := 0, len(nums) - 1

	for r > l {
		mid := l + (r-l) / 2

		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	var start, end int
	
	if target > nums[len(nums) - 1] {
		end = l - 1
		start = 0
	} else {
		start = l
		end = len(nums) - 1
	}

	for start <= end {
		mid := start + (end - start) / 2

		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			return mid
		}
	}

	return -1 
}
