func twoSum(nums []int, target int) []int {
	res := map[int]int{}

    for i := 0; i < len(nums); i++ {
		search :=  target - nums[i]
		v, ok := res[search]

		if ok && i != v {
			if i > v {
				return []int{v, i}
			}
			return []int{i, v}
		}

		res[nums[i]] = i
	}
	return []int{}
}
