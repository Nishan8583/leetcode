/*
1899. Merge Triplets to Form Target Triplet
Solved
Medium
Topics
Companies
Hint

A triplet is an array of three integers. You are given a 2D integer array triplets, where triplets[i] = [ai, bi, ci] describes the ith triplet. You are also given an integer array target = [x, y, z] that describes the triplet you want to obtain.

To obtain target, you may apply the following operation on triplets any number of times (possibly zero):

    Choose two indices (0-indexed) i and j (i != j) and update triplets[j] to become [max(ai, aj), max(bi, bj), max(ci, cj)].
        For example, if triplets[i] = [2, 5, 3] and triplets[j] = [1, 7, 5], triplets[j] will be updated to [max(2, 1), max(5, 7), max(3, 5)] = [2, 7, 5].

Return true if it is possible to obtain the target triplet [x, y, z] as an element of triplets, or false otherwise.



Example 1:

Input: triplets = [[2,5,3],[1,8,4],[1,7,5]], target = [2,7,5]
Output: true
Explanation: Perform the following operations:
- Choose the first and last triplets [[2,5,3],[1,8,4],[1,7,5]]. Update the last triplet to be [max(2,1), max(5,7), max(3,5)] = [2,7,5]. triplets = [[2,5,3],[1,8,4],[2,7,5]]
The target triplet [2,7,5] is now an element of triplets.

Example 2:

Input: triplets = [[3,4,5],[4,5,6]], target = [3,2,5]
Output: false
Explanation: It is impossible to have [3,2,5] as an element because there is no 2 in any of the triplets.

Example 3:

Input: triplets = [[2,5,3],[2,3,4],[1,2,5],[5,2,3]], target = [5,5,5]
Output: true
Explanation: Perform the following operations:
- Choose the first and third triplets [[2,5,3],[2,3,4],[1,2,5],[5,2,3]]. Update the third triplet to be [max(2,1), max(5,2), max(3,5)] = [2,5,5]. triplets = [[2,5,3],[2,3,4],[2,5,5],[5,2,3]].
- Choose the third and fourth triplets [[2,5,3],[2,3,4],[2,5,5],[5,2,3]]. Update the fourth triplet to be [max(2,5), max(5,2), max(5,3)] = [5,5,5]. triplets = [[2,5,3],[2,3,4],[2,5,5],[5,5,5]].
The target triplet [5,5,5] is now an element of triplets.



Constraints:

    1 <= triplets.length <= 105
    triplets[i].length == target.length == 3
    1 <= ai, bi, ci, x, y, z <= 1000



Solution:

Basically check if value=max(array1,array2), and then do same operation on result for rest of the values in array.
Question wants us to find if doing this will get us the target value.

Soln:
1. Remove any array with any value that is greater than values in target array, becaus max operation will result in value greater than the target
Example: [[2,5,3],[1,8,4],[1,7,5]], Remove [1,8,4], becaus even if doing max operation on 0 and 1 index, will have first positon value 2, the second position value will always result in 8
due to our max operation, and it will never reach our target of [2,7,5]

2. Second, now if the values are lower, then check if any of the values in the correct index matched, if it does add it to the set.
Ex: [2,5,3] 2 matchec, 5 does not (not the right index), 3 does not, add 2 to our set
[1,7,5] 7 and 5 both match.

Our set in the end will have [2,7,5] means some combination of array will result in the target

*/

package main

func mergeTriplets(triplets [][]int, target []int) bool {
	matched_values := []int{}

	// preventing adding the same position twice
	pos_added := map[int]int{}
	for _, value := range triplets {

		// if any of these values are greater than target, it means max operation would result in
		// this value being included in the result, so even if one matches, the other unnecessary greater
		// value would be included, which would not let us get the target value
		if value[0] > target[0] || value[1] > target[1] || value[2] > target[2] {
			continue
		}

		// fmt.Println("Checking",value)
		// loop, if value is same as target then add
		for i, v := range value {
			if v == target[i] {
				if _, ok := pos_added[i]; ok {
					continue
				}
				pos_added[i] = 1
				// fmt.Println("Something matched",v,i)
				matched_values = append(matched_values, v)
			}
			if len(matched_values) == 3 {
				return true
			}
		}
	}

	// fmt.Println("Final",matched_values)
	// by some combination we get that
	return len(matched_values) == 3
}
