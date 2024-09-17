/*
https://leetcode.com/problems/trapping-rain-water/description/
https://www.youtube.com/watch?v=ZI2z5pq0TqA&ab_channel=NeetCode
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.



Example 1:

Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.

formualae is min(max_left_value,max_righ_value)-height[i], so basically, the idea is get the lowest value between left and reight, cause thats
bottleneck where water can rise up to, then substract that with current height
lets say left is 4, right is 3, and current height is 1, then bottleneck here is 3, and water level can rise up to 3-1 = upto 2 level

SO how can we do this?
2 arrays one holds the most minimum value in the left for that index, and another holds the max value for that index
				 [0,1,0,2,1,0,1,3,2,1,2,1]
left max value = [0,0,1,1,2,2,2,2,3,3,3,3]
right max value= [3,3,3,3,3,3,3,3,2,2,1,0]
min(l,r) for index 0, its 0, substract 0-1 = -1 not possible and so on, sum it up

FOr 2 pointers, its a bit difficult
L=0, r= len(hjeight)-1
max_l and max_r, get min of both and substract by current value, if l is lower, increaser left, update max_l, if r is lower decrease right, update right
right max vaue =
*/
func trap(height []int) int {

	l := 0
	r := len(height) - 1
	max_l := height[l]
	max_r := height[r]
	sum := 0

	for {
		if l >= r {
			break
		}

		if max_l < max_r {
			l += 1
			if max_l < height[l] {
				max_l = height[l]
			}
			sum += max_l - height[l]
		} else {
			r -= 1
			if max_r < height[r] {
				max_r = height[r]
			}
			sum += max_r - height[r]
		}

	}
	return sum
}
