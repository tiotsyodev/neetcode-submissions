func productExceptSelf(nums []int) []int {
	prefmass := make([]int, len(nums))
	suffmass := make([]int, len(nums))
	pref := 1
	suf := 1

	res := []int{}

	for k, v := range nums {
		prefmass[k] = pref
		pref = v * pref
	}

	for i := len(nums) - 1; i >= 0; i-- {
		suffmass[i] = suf
		suf = suf * nums[i]
	}

	i := 0

	for range nums {
		res = append(res, prefmass[i]*suffmass[i])

		i++
	}

	return res
}
