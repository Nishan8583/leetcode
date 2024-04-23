// The question is a bit confusing
// Basically its asking to find max area between two indices, where height = minimun of value at both indices
// and width = distance between two indices
// left and right pointer, left at first, and right at last
// caculate area with above conditions, update max product if greater
// if value at left is higher, no point in changing it, so decrease r
// if value at right is higher, increase l
// keep on looping
// question: https://leetcode.com/problems/container-with-most-water/
// explaination: https://www.youtube.com/watch?v=UuiTKBwPgAo&ab_channel=NeetCode
// Damn this question was confusing man,
func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	max_area := 0
	for {
		if l > r {
			break
		}
		current_height := height[l]
		if height[r] < current_height { // bottleneck, choose the lowest height
			current_height = height[r]
		}

		width := r - l

		area := current_height * width
		if area > max_area {
			max_area = area
		}

		if height[l] < height[r] {
			l = l + 1
		} else {
			r = r - 1
		}
	}

	return max_area

}