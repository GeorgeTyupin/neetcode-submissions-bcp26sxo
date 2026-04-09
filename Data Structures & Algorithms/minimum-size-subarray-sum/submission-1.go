func minSubArrayLen(target int, nums []int) int {
	l, r := 0, -1
	ans := 0
	curLen, curSum := 0, 0

	for l < len(nums) {
		for r+1 < len(nums) && (curSum < target) {
			r++
			curSum += nums[r]
			curLen++
		}

		if ans == 0 && curSum >= target {
			ans = curLen
		} else if curSum >= target {
			ans = min(ans, curLen)
		}

		curSum -= nums[l]
		curLen--
		l++
	}

	return ans
}