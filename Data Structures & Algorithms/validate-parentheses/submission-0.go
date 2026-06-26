func isValid(s string) bool {
	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []rune{}
	strslice := []rune(s)

	for _, v := range strslice {
		if strings.ContainsRune("[{(", v) {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}

			if stack[len(stack) - 1] == brackets[v] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}