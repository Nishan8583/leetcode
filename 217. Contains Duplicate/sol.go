// 217. Contains Duplicate
// loop through, check in map, if present return true, else
// append in map
// https://leetcode.com/problems/contains-duplicate/
func containsDuplicate(nums []int) bool {
	appeared := map[int]struct{}{}

	for _, v := range nums {
		if _, ok := appeared[v]; ok {
			return true
		}
		appeared[v] = struct{}{}
	}

	return false
}

// Sorting, is more efficient
import "sort"

func containsDuplicate(nums []int) bool {
    sort.Ints(nums)

    for i:= 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            return true
        }
    }
    return false 
}