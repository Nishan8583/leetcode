
/*
Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
https://leetcode.com/problems/top-k-frequent-elements/description/
https://www.youtube.com/watch?v=YPTqKIgVk-k&t=1s&ab_channel=NeetCode
first create a map that holds number to count
an arrray of length len(nums)+1, treat index as count, and value is array of elements having that count
loop through, from last to first, cause we need elements with most values, decrease k
Observation, treating index of array with some meaning, in previous question index was an array of element occurence and now count
this is interesting.
*/
func topKFrequent(nums []int, k int) []int {
	num_to_count := make(map[int]int)
	freq := make([][]int, len(nums)+1) // index is count, and value is going to be a slice of values that have that count
	result := make([]int, 0, k)
	for _, v := range nums {
		if count, ok := num_to_count[v]; ok {
			count += 1
			num_to_count[v] = count
		} else {
			num_to_count[v] = 1
		}
	}

	// append in array
	for num, count := range num_to_count {

		freq[count] = append(freq[count], num)
	}

	for i := (len(freq) - 1); i >= 0; i-- {
		if len(freq[i]) > 0 {
			for _, value := range freq[i] {
				if k == 0 {
					break
				}
				result = append(result, value)
				k -= 1

			}
		}
	}
	return result

}