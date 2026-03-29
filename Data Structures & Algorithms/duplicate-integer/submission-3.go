func hasDuplicate(nums []int) bool {
    seen := make(map[int]struct{})

	for _, x := range nums {
		if _, ok := seen[x]; ok {
			return true
		}
		seen[x] = struct{}{}
	}

	return false
}
