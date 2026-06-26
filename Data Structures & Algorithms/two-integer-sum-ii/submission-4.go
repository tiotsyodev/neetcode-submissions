func twoSum(numbers []int, target int) []int {
	small := 0
	big := len(numbers) - 1

	for small < big {
		sum := numbers[small] + numbers[big]

		if sum > target {
			big--
		} else if sum < target {
			small++
		} else {
			return []int{small + 1, big + 1}
		}
	}

	return []int{0, 0}
}
