// Since the array is sorted, even if we do a rotation, we just 
// need to check if the current value is less than earlier value i.e. value on the left
// if it is, we have found the start of the original array, and thus the minimum value.
func findMin(nums []int) int {
    current_value := 0   // useful in cases where the length of array is just one
    for i,v := range nums {
        current_value = v  // updating current value
        index := i
        if i ==0 {  // if its the first index, wrap around to the last
            index = len(nums)-1 
        } else {
            index = index - 1
        }
        if v < nums[index] {
            return v
        }

    }
    return current_value
}
