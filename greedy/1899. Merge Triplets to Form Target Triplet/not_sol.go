package main

import "fmt"

func mergeTriplets(triplets [][]int, target []int) bool {
	// take the first triplet, since there is no other values before it, consider this as max
	for i := 0; i < len(triplets); i++ {

		max_int := triplets[i]

		// initially count is going to be 0
		count := 0

		// loop through target, check if values is same as target,
		// increase count on match
		for k, v := range target {
			if v == max_int[k] {
				count++
			}
		}

		fmt.Println("Beginning count is", count)
		// if count is 3, we got our target
		if count == 3 {
			return true
		}

		// Now loop through all our triplets, starting from position one
		for j := 0; j < len(triplets); j++ {

			// Perform the max operation as mentioned by the question and get the result
			temp_slice := maxOperation(max_int, triplets[j])
			fmt.Printf(
				"Output of max operation for max_int %v and %v is %v\n",
				max_int,
				triplets[j],
				temp_slice,
			)
			// holds the matched count temporarily
			temp_count := 0

			// check the count of matched with target
			for k, v := range target {
				if v == temp_slice[k] {
					temp_count++
				}
			}

			// if current received value from max operation is more than previous matched count, replace the max_int with this new one
			if temp_count > count {
				fmt.Println("found where count is greater ", max_int, temp_count, count)
				max_int = temp_slice
				count = temp_count
			}

			// we could have found the match inside our loop, no need to proceed further
			if count == 3 {
				return true
			}
		}

		fmt.Println("Final Count is", count)
		// if count is three, we have found our match
		if count == 3 {
			return true
		}
	}

	// we reach here, when we do not find the target
	return false
}

func maxOperation(slice1, slice2 []int) []int {
	maxSlice := make([]int, 3)
	for i, v := range slice1 {
		if v >= slice2[i] {
			maxSlice[i] = v
		} else {
			maxSlice[i] = slice2[i]
		}
	}

	return maxSlice
}

func main() {
	input := [][]int{
		{7, 15, 15},
		{11, 8, 3},
		{5, 3, 4},
		{12, 9, 9},
		{5, 12, 10},
		{7, 15, 10},
		{7, 6, 4},
		{3, 9, 8},
		{2, 13, 1},
		{14, 2, 3},
	}

	fmt.Println(mergeTriplets(input, []int{14, 6, 4}))
}
