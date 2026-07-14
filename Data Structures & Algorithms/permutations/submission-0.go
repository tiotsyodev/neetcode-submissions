func permute(nums []int) [][]int {
    res := [][]int{}
    path := make([]bool, len(nums))
    sub := []int{}

    var btr func()
    btr = func() {

        if len(sub) == len(nums) {
            temp := make([]int, len(sub))
            copy(temp, sub)
            res = append(res, temp)
            return
        }

        for i := 0; i < len(nums); i++ {
            if path[i] {
                continue
            }
            sub = append(sub, nums[i])
            path[i] = true
            btr()
            path[i] = false
            sub = sub[:len(sub)-1]
        }
    }

    btr()

    return res
}
