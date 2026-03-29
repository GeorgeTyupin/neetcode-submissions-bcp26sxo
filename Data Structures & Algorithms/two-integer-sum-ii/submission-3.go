func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for l < r {
		if numbers[r] + numbers[l] == target {
			return []int{l+1, r+1}
		} else if numbers[r] + numbers[l] > target {
			r--
		} else {
			l++
		}
	}

	return []int{-1, -1}
}
