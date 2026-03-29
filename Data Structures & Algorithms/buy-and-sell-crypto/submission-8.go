func maxProfit(prices []int) int {
	l, result := 0, 0

	for r := 1; r < len(prices); r++ {
		profit := prices[r] - prices[l]

		if profit < 0 {
			l = r
			continue
		}
		if profit > result {
			result = profit
		}
	}

	return result
}
