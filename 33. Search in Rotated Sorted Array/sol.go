func search(nums []int, target int) int {
    for i,v := range nums {
        if target == v {
            return i
        } 
        if target < v{
            // its less than the current index value, so lets rotate backwards
            if target < v {
                local_index := len(nums) - 1
                for j:=local_index;j >= 0; j-- {
                    if  target == nums[local_index] {
                        return j
                    }

                    // if target is less than current value, we have reached a point 
                    // sorted array where values are greater than target
                    if target > nums[local_index]  {
                        return -1
                    }
                }
            }
        }

    }
    return -1
}
