func combinationSum2(candidates []int, target int) [][]int {
    res := [][]int{}
    cursubset := []int{}
    sum := 0

    sort.Ints(candidates)

    var backtrack func(idx int)
    backtrack = func(idx int) {

        if sum > target {
            return
        }

        if sum == target {
            subset := make([]int, len(cursubset))
            copy(subset, cursubset)
            res = append(res, subset)
            return
        }

        for i := idx; i < len(candidates); i++ {
            if i > idx && candidates[i - 1] == candidates[i] {
                continue
            }

            sum += candidates[i]
            cursubset = append(cursubset, candidates[i])

            backtrack(i + 1)

            sum -= candidates[i]
            cursubset = cursubset[:len(cursubset)-1]
        }

    } 

    backtrack(0)

    return res
}
