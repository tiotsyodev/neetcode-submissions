func groupAnagrams(strs []string) [][]string {
    resmap := map[[27]int][]string{}

    for _, v := range strs {
        var mass [27]int
        str := []rune(v)
        for _, letter := range str {
            pos := letter - 'a' + 1
            mass[pos] = mass[pos] + 1
        }
       resmap[mass] = append(resmap[mass], v)
        
    }

    res := [][]string{}
    for _, value := range resmap{
        res = append(res, value)
    }
    return res

}
