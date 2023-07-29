// https://leetcode.com/problems/search-in-rotated-sorted-array/submissions/
//https://youtu.be/U8XENwh8Oy8
// try to find in which part of array we are, cause even if the array is sorted, we can determine which
// side of array we are on, and then update Left and right accordingly
func search(nums []int, target int) int {
    l,r := 0,len(nums)-1
    

    for {
        mid := int((l+r)/2)
       

        mid_value := nums[mid]
        
        if mid_value == target {
            return mid
        }
        if l >= r {
            break
        }
        // is it in left sorted array?
        if nums[l] <= mid_value {
            if (target  > mid_value) || (target < nums[l]) {
                // value won't be on the left sorted array
                l = mid + 1
            } else {
                // value migh be on the left, so search left
                r = mid - 1
            }
        } else {
            // we are in right sorted array
            if (target < mid_value) || (target > nums[r]) {
                r = mid -1
            } else {
                l = mid +1
            }
        }
        
    }
    return -1
}
