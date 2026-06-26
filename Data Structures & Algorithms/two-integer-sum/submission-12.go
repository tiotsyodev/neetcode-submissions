func twoSum(nums []int, target int) []int {
	res := map[int]int{}

    for i := 0; i < len(nums); i++ {
		search :=  target - nums[i]
		if v, ok := res[search]; ok {
			return []int{v, i}
		}

		res[nums[i]] = i
	}
	return []int{}
}
