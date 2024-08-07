/*

846. Hand of Straights
Solved
Medium
Topics
Companies

Alice has some number of cards and she wants to rearrange the cards into groups so that each group is of size groupSize, and consists of groupSize consecutive cards.

Given an integer array hand where hand[i] is the value written on the ith card and an integer groupSize, return true if she can rearrange the cards, or false otherwise.



Example 1:

Input: hand = [1,2,3,6,2,3,4,7,8], groupSize = 3
Output: true
Explanation: Alice's hand can be rearranged as [1,2,3],[2,3,4],[6,7,8]

Example 2:

Input: hand = [1,2,3,4,5], groupSize = 4
Output: false
Explanation: Alice's hand can not be rearranged into groups of 4.


Solution:
Sorted array
count of elelemtans,
wierd shit here, look at youtube video
*/

package main

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	countMap := make(map[int]int)

	for _, num := range hand {
		countMap[num]++
	}

	var minHeap []int

	for uniqueNum := range countMap {
		minHeap = append(minHeap, uniqueNum)
	}

	sort.Ints(minHeap)

	for len(minHeap) != 0 {
		first := minHeap[0]

		for i := first; i < groupSize+first; i++ {
			if _, ok := countMap[i]; !ok {
				return false
			}

			countMap[i]--
			val := countMap[i]

			if val == 0 {
				if i != minHeap[0] {
					return false
				}

				minHeap = minHeap[1:]
			}

		}

	}
	return true
}
