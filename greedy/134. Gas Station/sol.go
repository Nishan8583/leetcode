/*
There are n gas stations along a circular route, where the amount of gas at the ith station is gas[i].

You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from the ith station to its next (i + 1)th station. You begin the journey with an empty tank at one of the gas stations.

Given two integer arrays gas and cost, return the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return -1. If there exists a solution, it is guaranteed to be unique



Example 1:

Input: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
Output: 3
Explanation:
Start at station 3 (index 3) and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 4. Your tank = 4 - 1 + 5 = 8
Travel to station 0. Your tank = 8 - 2 + 1 = 7
Travel to station 1. Your tank = 7 - 3 + 2 = 6
Travel to station 2. Your tank = 6 - 4 + 3 = 5
Travel to station 3. The cost is 5. Your gas is just enough to travel back to station 3.
Therefore, return 3 as the starting index.

Example 2:

Input: gas = [2,3,4], cost = [3,4,3]
Output: -1
Explanation:
You can't start at station 0 or 1, as there is not enough gas to travel to the next station.
Let's start at station 2 and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 0. Your tank = 4 - 3 + 2 = 3
Travel to station 1. Your tank = 3 - 3 + 3 = 3
You cannot travel back to station 2, as it requires 4 unit of gas but you only have 3.
Therefore, you can't travel around the circuit once no matter where you start.const
Solution:
gas = [1,2,3,4,5]
cost = [3,4,5,1,2]
1. First we check if sum of all elements in gas is greater or equals to costSUm, if not, it means that the gas is neven going to be enough, and return -1.
2. Create a diff array, each element will basically be the cost of gas for jumping to the next step, that way in the second loop, we can just add the diff elmeent i.e. increase the gas/collect gas
diff becomse [-2,-2,-2,3,3]
3. Find the element thats non negative
4. Check if we can reach the final position from that position
5. no need to loop to the front, since we have already figured out that sumGas > sumCost, so if it reached the end, it can complete the cycle
*/

package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	sumGas := 0
	sumCost := 0
	startIndex := -1

	diff := make([]int, len(gas))

	for i := 0; i < len(gas); i++ {
		sumGas += gas[i]
		sumCost += cost[i]
		diff[i] = gas[i] - cost[i]

		// starting position have not been found yet
		if startIndex == -1 {
			if diff[i] >= 0 { // since its gte 0, we can jump to next, so this could be our staring position
				startIndex = i
			}
		}
	}

	fmt.Println("Diff", diff)
	// if the total sum of cost is greater than sumGas, we can not complete a loop
	if sumCost > sumGas {
		return -1
	}

	// no value gte 0 was found
	if startIndex == -1 {
		return startIndex
	}

	totalGas := 0
	for i := startIndex; i < len(diff); i++ {
		totalGas += diff[i]

		if totalGas < 0 {
			totalGas = 0
			startIndex = i + 1
		}
	}
	return startIndex
}

func main() {
	fmt.Println(canCompleteCircuit([]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}))
}
