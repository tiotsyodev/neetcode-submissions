func findDuplicate(nums []int) int {
	f := nums[0]
	s :=  nums[0]

	for {
		s = nums[s]
		f = nums[nums[f]]

		if f == s {
			break
		}
	}

	s = nums[0]

	for s != f {
		s = nums[s]
		f = nums[f]
	}

	return s
}
