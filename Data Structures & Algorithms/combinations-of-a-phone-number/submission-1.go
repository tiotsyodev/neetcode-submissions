func letterCombinations(digits string) []string {
    if digits == "" {
        return []string{}
    }
	
    digitsmap := map[byte]string{
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }

    res := []string{}
    cur := []rune{}

    var dfs func(idx int)
    dfs = func(idx int) {
        if idx == len(digits) {
            res = append(res, string(cur))
            return
        }

        for _, ch := range digitsmap[digits[idx]] {
            cur = append(cur, ch)
            dfs(idx + 1)
            cur = cur[:len(cur)-1]
        }
    }

    dfs(0)

    return res
}
