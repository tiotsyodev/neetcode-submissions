func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		l, r := i + 1, len(nums) - 1

		for r > l {
			if nums[l] + nums[r] + nums[i] > 0 {
				r--
			} else if nums[l] + nums[r] + nums[i] < 0 {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for r > l && nums[r] == nums[r - 1] {
					r--
				}
				for r > l && nums[l] == nums[l + 1] {
					l++
				}
				r--
				l++
			}
		}
	}

	return res

}
