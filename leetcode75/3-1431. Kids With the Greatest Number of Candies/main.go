package main

import "slices"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	greatest := slices.Max(candies)
	result := make([]bool, len(candies))

	for index, value := range candies {
		if value+extraCandies >= greatest {
			result[index] = true
		}
	}
	return result
}
