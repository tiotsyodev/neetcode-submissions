func topKFrequent(nums []int, k int) []int {
	resmap := map[int]int{}
	res := []int{}

	for _, val := range nums {
		resmap[val] += 1
	}

	freq := make([][]int, len(nums) + 1)

	for k, v := range resmap {
		freq[v] = append(freq[v], k)
	}

	for i := len(freq) - 1; i > 0; i-- {
		for _, v := range freq[i] {
			res = append(res, v)

			if len(res) == k {
				return res
			}
		}
	}

	return res
}
