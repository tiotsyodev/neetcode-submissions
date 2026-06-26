func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	windowlen := len(s1)
	substringms := [26]int{}
	stringms := [26]int{}
	for _, v := range s1 {
		substringms[v - 'a']++
	}

	for i := 0; i < windowlen; i++ {
		stringms[s2[i] - 'a']++
	}

	if substringms == stringms {
		return true
	}

	for r := windowlen; r < len(s2); r++ {
		stringms[s2[r] - 'a']++
		stringms[s2[r - windowlen] - 'a']--

		if substringms == stringms {
			return true
		}
	}

	return false
}
