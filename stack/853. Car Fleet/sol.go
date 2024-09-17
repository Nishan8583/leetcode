/*
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

import "fmt"

// stack, each position will hold the number of cars at that position
type stack []int

func newStack(size int) stack {
	s := make([]int, size)
	return s
}

// get will return the value at the particular index, in this the case the number of cars
func (s *stack) get(index int) int {
	return (*s)[index]
}

// push the car by a certain margin, in this case, drive the car by certain miles
func (s *stack) push(car, cur_index, new_index int) {
	temp := (*s)[new_index]    // temporarily hold the value that is currently in new index
	(*s)[new_index] = temp + 1 // increase the number of the cars in the desired position

	// -1 means the stack is initializing
	if cur_index == -1 {
		return
	}
	(*s)[cur_index] -= (*s)[new_index] - 1 // also remove the car from the current position that it has
}

// decrease the number of cars
func (s *stack) pop(car, index int) {
	if (*s)[index] == 0 {
		return
	}
	(*s)[index] = (*s)[index] - 1 // remove an element
}

// returns the number of targets that were presetn int the target in stack
func (s *stack) getNumberOfCarsAtTarget() int {
	value := (*s)[len(*s)-1]
	(*s)[len(*s)-1] = 0
	return value
}
func carFleet(target int, position []int, speed []int) int {

	carFleet := 0 // the number of car fleets that have completed the journey

	// initializing the stack with the values
	s := newStack(target + 1)
	for i := 0; i <= target; i++ {
		s.push(i, -1, 1)
	}
	fmt.Println(s)
	c := 0 // the number of car fleets that have finished the journey
	for {
		if c == len(position) {
			break
		}

		// travel
		for i, v := range speed {
			cur_pos := position[i]
			if cur_pos >= target {
				continue
			}
			position[i] += v
			s.push(i, cur_pos, position[i])
		}
		temp := s.getNumberOfCarsAtTarget()
		if temp != 0 {
			fmt.Println("GOt numbebr of cars", temp)
			fmt.Println(position)
			c += temp
			carFleet += 1
		}

	}

	return carFleet
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
	}

	for _, tests := range testCases {
		r := carFleet(tests.target, tests.position, tests.speed)
		if r != tests.eCarFleet {
			fmt.Printf("TEST CASE FAILED expected %d got %d \n", tests.eCarFleet, r)
		}
	}
}
