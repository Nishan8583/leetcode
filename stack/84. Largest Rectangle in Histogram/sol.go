// https://leetcode.com/problems/largest-rectangle-in-histogram/submissions/1232337186/
// https://www.youtube.com/watch?v=zx5Sw9130L0

package main

import "fmt"

// We are gooing to go from left to right
// if previous height was bigger than current, we need to pop that off
// since the current one is smaller, our index can be decreased since we can extend back  too
// watch the video, this is a bit confusing
func largestRectangleArea(heights []int) int {
	max_area := 0
	stack := [][]int{}

	for i, h := range heights {
		start := i

		// if the value at top of the stack is greater than current value
		// then the previous value can not be increased
		for len(stack) != 0 && stack[len(stack)-1][1] > h {
			ihSlice := stack[len(stack)-1]
			index := ihSlice[0]
			height := ihSlice[1]
			stack = stack[:len(stack)-1]
			start = index // since the previous rectangle is > ,the current rectangle can be increased to the right
			area := height * (i - index)
			if area > max_area {
				max_area = area
			}
		}
		stack = append(stack, []int{start, h})
	}

	for _, ihSlice := range stack {
		i := ihSlice[0]
		h := ihSlice[1]
		area := h * (len(heights) - i)
		if area > max_area {
			max_area = area
		}
	}

	return max_area
}

func main() {
	input := []int{2, 1, 5, 6, 2, 3}
	//fmt.Println(largestRectangleArea(input))
	//input := []int{0, 9}
	fmt.Println(largestRectangleArea(input))

}
