func twoSum(nums []int, target int) []int {
	res := map[int]int{}

    for i := 0; i < len(nums); i++ {
		if v, ok := res[target - nums[i]]; ok {
			return []int{v, i}
		}

		res[nums[i]] = i
	}
	return []int{}
}
