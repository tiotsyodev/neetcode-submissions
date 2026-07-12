func partition(s string) [][]string {
    res := [][]string{}
    subms := []string{}

    var btr func(start int) 
    btr = func(start int) {

        if start == len(s) {
            temp := append([]string{}, subms...)
            res = append(res, temp)
            return
        }


        for i := start; i < len(s); i++ {
            substr := s[start:i+1]

            if isPalindrome(substr) {
                subms = append(subms, substr)
                btr(i + 1)
                subms = subms[:len(subms)-1]
            }
        }
    
    }

    btr(0)

    return res
}


func isPalindrome(s string) bool {
    l := 0
    r := len(s) - 1

    for l < r {
        if s[l] != s[r] {
            return false
        }
        l++
        r--
    }

    return true
}