func hasDuplicate(nums []int) bool {
    dups := map[int]int{}

    for _, v := range nums {
        _, ok := dups[v]
        if !ok {
            dups[v] = 1
        } else {
            return true
        }
    }

    return false
}
