
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
// https://www.youtube.com/watch?v=1pkOgXD63yU&ab_channel=NeetCode
// Use sliding windows
// hold left pointer and right pointer, Lr = buyvalue, Rr = sell value
// Rr should Lr should always be less then sell value
// If not update buy value
// Check profit with current Rr - Lr
// Keep updating the Rr, as we move to the right
func maxProfit(prices []int) int {
	max_profit := 0

	if len(prices) == 0 {
		return 0
	}

	buy_value := prices[0]

	for i, v := range prices {
		curr_buy_value := v
		if (i + 1) >= len(prices) {
			break
		}
		curr_sell_value := prices[i+1]

		// if tomorrows sell value is less then our chosen buy value
		// t hen we are in loss, fix it
		if buy_value > curr_buy_value {
			buy_value = curr_buy_value
		}

		curr_profit := curr_sell_value - buy_value
		if curr_profit >= max_profit {
			max_profit = curr_profit
		}
	}

	return max_profit
}