type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	restr := ""
	for _, v := range strs {
		restr += strconv.Itoa(len([]rune(v))) + "#" + v
	}

	return restr
}

func (s *Solution) Decode(encoded string) []string {
	var res []string
	runes := []rune(encoded)
	i := 0

	for i < len(runes) {
		j := i
		for runes[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(string(runes[i:j]))

		startOfStr := j + 1
		endOfStr := startOfStr + length
		res = append(res, string(runes[startOfStr:endOfStr]))

		i = endOfStr
	}

	return res
}