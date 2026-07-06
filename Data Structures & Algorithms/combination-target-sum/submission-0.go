func combinationSum(nums []int, target int) [][]int {
    res := [][]int{}
    cur := 0
    cursubset := []int{}
    
    var backtrack func(start int)
    backtrack = func(start int) {

        if cur == target {
            subset := make([]int, len(cursubset))
            copy(subset, cursubset)
            res = append(res, subset)
            return
        }


        if cur > target {
            return
        }

        for i := start; i < len(nums); i++ {
            cur += nums[i]
            cursubset = append(cursubset, nums[i])

            backtrack(i)

            cur -= nums[i]
            cursubset = cursubset[:len(cursubset) - 1]
        }

    }

    backtrack(0)

    return res
}


