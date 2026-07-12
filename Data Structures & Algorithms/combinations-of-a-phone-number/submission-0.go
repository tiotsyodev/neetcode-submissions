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
    cur := make([]byte, 0, len(digits))

    var btr func(idx int)
    btr = func(idx int) {
        if idx == len(digits) {
            res = append(res, string(cur))
            return
        }

        for _, v := range digitsmap[digits[idx]] {
            cur = append(cur, byte(v))
            btr(idx + 1)
            cur = cur[:len(cur) - 1]
        }
    }

    btr(0)

    return res
}
