func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	subset := []int{}

	sort.Ints(nums)

	var backtrack func(idx int)
	backtrack = func(idx int) {

		cursubset := make([]int, len(subset))
		copy(cursubset, subset)
		res = append(res, cursubset)

		for i := idx; i < len(nums); i++ {
			if i > idx && nums[i] == nums[i - 1] {
				continue
			}

			subset = append(subset, nums[i])

			backtrack(i + 1)

			subset = subset[:len(subset) - 1]
		}

	}

	backtrack(0)

	return res	
}
