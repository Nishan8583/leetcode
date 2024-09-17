package main

import "fmt"

// DYnamic programming
// Save the count of coins for a sum
// DP[0] = requires 0 coins
// DP[1] = min(dp[1],1+dp[1-current_coin])// 1(because we need to add one count for this) + how much coin required for previous count
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i, _ := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i < amount+1; i++ {

		for _, c := range coins {
			if i-c >= 0 {
				fmt.Println("found", i, c)
				dp[i] = min(dp[i], 1+dp[i-c])
			}
		}
	}
	if dp[amount] >= amount+1 {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 100))
}
