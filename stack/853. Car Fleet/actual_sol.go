/*
https://leetcode.com/problems/car-fleet/description/
neetcode: https://youtu.be/Pr6T-3yB9RM?si=JbHPVZvJ2S-0NMUD
There are n cars going to the same destination along a one-lane road. The destination is target miles away.

You are given two integer array position and speed, both of length n, where position[i] is the position of the ith car and speed[i] is the speed of the ith car (in miles per hour).

A car can never pass another car ahead of it, but it can catch up to it and drive bumper to bumper at the same speed. The faster car will slow down to match the slower car's speed. The distance between these two cars is ignored (i.e., they are assumed to have the same position).

A car fleet is some non-empty set of cars driving at the same position and same speed. Note that a single car is also a car fleet.

If a car catches up to a car fleet right at the destination point, it will still be considered as one car fleet.

Return the number of car fleets that will arrive at the destination.

Example 1:

Input: target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
Output: 3
Explanation:
The cars starting at 10 (speed 2) and 8 (speed 4) become a fleet, meeting each other at 12.
The car starting at 0 does not catch up to any other car, so it is a fleet by itself.
The cars starting at 5 (speed 1) and 3 (speed 3) become a fleet, meeting each other at 6. The fleet moves at speed 1 until it reaches target.
Note that no other cars meet these fleets before the destination, so the answer is 3.
Example 2:

Input: target = 10, position = [3], speed = [3]
Output: 1
Explanation: There is only one car, hence there is only one fleet.
Example 3:

Input: target = 100, position = [0,2,4], speed = [4,2,1]
Output: 1
Explanation:
The cars starting at 0 (speed 4) and 2 (speed 2) become a fleet, meeting each other at 4. The fleet moves at speed 2.
Then, the fleet (speed 2) and the car starting at 4 (speed 1) become one fleet, meeting each other at 6. The fleet moves at speed 1 until it reaches target.
*/

package main

import (
	"fmt"
	"sort"
)

// We basically create a set of position and speed
// Sort according to position
// Calculate time taken to reach destination in reverse order
// put it in stack,
// continue, if the new car can reach destination at same time or early, then they reach destination together
// so they are same fleet
// remove the new car, because they are the same fleet
// return length
func carFleet(target int, position []int, speed []int) int {
	set := [][]int{}
	for i := 0; i < len(position); i++ {
		set = append(set, []int{position[i], speed[i]})
	}

	// sorting according to first element in index
	sort.Slice(set, func(i, j int) bool {
		return set[i][0] < set[j][0]
	})

	stack := []float32{}

	for i := len(set) - 1; i >= 0; i-- {
		time_taken := float32((float32(target) - float32(set[i][0])) / float32(set[i][1]))
		//fmt.Printf("Time taken for pos %d with speed %d is %f \n", set[i][0], set[i][1], time_taken)
		stack = append(stack, time_taken)
		if len(stack) >= 2 {

			// if the last elemen pushed, will reach earlier than second to last element
			// then they will reach destination at same speed, thus same feet
			if stack[len(stack)-1] <= stack[len(stack)-2] {

				stack = stack[:len(stack)-1] // poping the stack
			}
		}
	}
	fmt.Println(len(stack))
	return len(stack)
}

func main() {
	type test struct {
		position  []int
		speed     []int
		target    int
		eCarFleet int
	}

	testCases := []test{
		{position: []int{10, 8, 0, 5, 3}, speed: []int{2, 4, 1, 1, 3}, target: 12, eCarFleet: 3},
		{position: []int{6, 8}, speed: []int{3, 2}, target: 10, eCarFleet: 2},
	}

	for _, tests := range testCases {
		r := carFleet(tests.target, tests.position, tests.speed)
		if r != tests.eCarFleet {
			fmt.Printf("TEST CASE FAILED expected %d got %d \n", tests.eCarFleet, r)
		}
	}
}
