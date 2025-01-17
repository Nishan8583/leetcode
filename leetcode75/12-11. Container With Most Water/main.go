package main

import "fmt"

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	max_area := 0

	for left < right {
		min_height := min(height[left], height[right])
		area := min_height * (right - left)
		max_area = max(max_area, area)

		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
	}

	return max_area
}

func main() {
	fmt.Println([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
}
