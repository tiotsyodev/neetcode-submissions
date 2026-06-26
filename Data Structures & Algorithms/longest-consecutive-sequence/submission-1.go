func longestConsecutive(nums []int) int {
	set := map[int]bool{}

	for _, v := range nums {
		set[v] = true
	}
	mx := 0
	for v, _ := range set {
		if _, ok := set[v - 1]; !ok {
			counter := 1
			thisVal := v + 1
			for {
				if _, ok := set[thisVal]; ok {
					counter++
					thisVal++
				} else {
					break
				}
			}

			if mx < counter {
				mx = counter
			}
		}
	}

	return mx

}
