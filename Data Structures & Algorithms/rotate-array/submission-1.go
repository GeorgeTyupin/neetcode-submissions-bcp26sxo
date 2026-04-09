func rotate(nums []int, k int) {
	for range k {
		cur := nums[0]
		for p := range len(nums) - 1 {
			next := nums[p+1]
			nums[p+1] = cur
			cur = next
		}
		nums[0] = cur
	}
}
