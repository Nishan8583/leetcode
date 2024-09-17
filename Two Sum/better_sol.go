// https://leetcode.com/problems/two-sum/
// 1. Create a hashmap to hold value to index
// 2. Loop through the array
// 3. Substract the current value from target
// 4. Check if the result is present in hashmap
// 5. If not add to map
// Ex: [2,1,5,3] target 5
// first iteration hashmap is empty,4-2 = 2, value not present in hashmap append
// second iteration hashmap = {2:0}, 4-1 =3 value not present in hashmap append
// third iteration hashmap = {2:0,1:1}, 4-5 =-1 value not present in hashmap append
// fourth iteration hashmap = {2:0,1:1,5:2}, 4-3 =1 value present, return index 1 and 3's index
// with this we do not have to loop over array twice, since we have added value in hashmap
func twoSum(nums []int, target int) []int {
	// holds values to hashmap
	checkedValues := map[int]int{}

	indices := []int{-1, -1}
	for i, v1 := range nums {
		requiredValue := target - v1
		if index, ok := checkedValues[requiredValue]; ok {
			return []int{index, i}
		}
		checkedValues[v1] = i
	}
	return indices
}