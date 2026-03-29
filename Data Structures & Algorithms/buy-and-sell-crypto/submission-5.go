func maxProfit(prices []int) int {
    l, p := 0, 0
	for r := 1; r < len(prices); r++ {
		if prices[r] - prices[l] < 0 {
			l = r
			continue
		}
		if prices[r] - prices[l] > p {
			p = prices[r] - prices[l]
		}
	}

	return p
}
