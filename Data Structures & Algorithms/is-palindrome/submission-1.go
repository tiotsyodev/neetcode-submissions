func isPalindrome(s string) bool {
	sl := []rune(strings.ToLower(s))

	l := 0
	r := len(sl) - 1

	for r >= l {

		for l < r && !(unicode.IsLetter(sl[l]) || unicode.IsDigit(sl[l])) {
			l++
		}

		for l < r &&  !(unicode.IsLetter(sl[r]) || unicode.IsDigit(sl[r])) {
			r--
		}


		if sl[r] != sl[l] {
			return false
		}

		r--
		l++
	}

	return true
}
