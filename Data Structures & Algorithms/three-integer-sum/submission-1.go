
func threeSum(nums []int) [][]int {
	res := [][]int{}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
            continue
        }

		l := i + 1
		r := len(nums) - 1
		searchFor := -nums[i]

		for l < r {
			if nums[r] + nums[l] == searchFor {
				res = append(res, []int{nums[i], nums[r], nums[l]})
				for l < r && nums[l] == nums[l+1] {
                    l++
                }
                for l < r && nums[r] == nums[r-1] {
                    r--
                }
                
                r--
                l++
			} else if nums[r] + nums[l] > searchFor {
				r--
			} else if nums[r] + nums[l] < searchFor { 
				l++
			}
		}
	}
	
	return res
}