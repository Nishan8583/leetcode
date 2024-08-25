/*
746. Min Cost Climbing Stairs
Solved
Easy
Topics
Companies
Hint

You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you pay the cost, you can either climb one or two steps.

You can either start from the step with index 0, or the step with index 1.

Return the minimum cost to reach the top of the floor.

Example 1:

Input: cost = [10,15,20]
Output: 15
Explanation: You will start at index 1.
- Pay 15 and climb two steps to reach the top.
The total cost is 15.

Example 2:

Input: cost = [1,100,1,1,1,100,1,1,100,1]
Output: 6
Explanation: You will start at index 0.
- Pay 1 and climb two steps to reach index 2.
- Pay 1 and climb two steps to reach index 4.
- Pay 1 and climb two steps to reach index 6.
- Pay 1 and climb one step to reach index 7.
- Pay 1 and climb two steps to reach index 9.
- Pay 1 and climb one step to reach the top.
The total cost is 6.

Constraints:

	2 <= cost.length <= 1000
	0 <= cost[i] <= 999
*/

/*
Solve it with dynamic tree style

calculate from last
Watch this video
https://www.youtube.com/watch?v=ktmzAZWkEZ0
*/
package main

import "fmt"

func minCostClimbingStairs(cost []int) int {

	cost = append(cost, 0)
	// [10,15,20,0]
	// each position will now hold the minimum value it takes to reach to the end
	// for 15, one jump is 15+20, or two jump is 15+0 15, min is 15
	// Same for 10
	// Which one among 0th and 1th is smallest
	// return that
	// Since one element can jump only 1 or 2 steps
	// we caculate from last, the min value it takes to reach the end
	//One jump 35 two jump 15 at 1
	// i = len(cost)-3 because, each step max jump is 2
	// possible jump is cost[i] + cost[i+1] i.e. 10+15 = 25, or cost[i+2] 2 jump 15+20 = 35
	// update i with the min value
	// [10 15 20 0]
	// One jump 25 two jump 30 at 0
	// [25 15 20 0]
	for i := len(cost) - 3; i >= 0; i-- {
		oneJump := cost[i] + cost[i+1]
		twoJump := cost[i] + cost[i+2]
		fmt.Printf("One jump %d two jump %d at %d\n", oneJump, twoJump, i)

		// Update i to the min value it takes to reach the end
		if oneJump < twoJump {
			cost[i] = oneJump
		} else {
			cost[i] = twoJump
		}
		fmt.Println(cost)
	}

	if cost[0] < cost[1] {
		return cost[0]
	}
	return cost[1]
}

func main() {

	cost := []int{10, 15, 20}

	fmt.Println(minCostClimbingStairs(cost))
}
