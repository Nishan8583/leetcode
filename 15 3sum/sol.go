
// First sort the number, saince there is possibility of duplicates, sorting will help us quickly check if previous number is same as current number
// Then loop through array, work on left right pointer, left will hold current +1 and R will hold final value, add all of them 
// IF sum is greater than 0, move R to left, since we know left is lower since array is sorted now
// if sum is less than 0, move L to right, since we know right is bigger since array is sorted now
// Else means sum is 0, append it,
// now need to update L, so l+1, check if new is same as previous L, if so break, or also break if l>=r cause we are at the end
// I could not solve it myself obviously, https://youtu.be/jzZsG8n2R9A
// https://leetcode.com/problems/3sum/submissions/
package main
import (
  "fmt"
  "sort"
)


func threeSum(nums []int) [][]int {
    res := [][]int{}  // whill hold the result
    sort.Ints(nums)  // sorting, so that we can check for duplicates

    for i,a := range nums {
        // avoid duplicates
        if (i > 0) && (a == nums[i-1]) {
            continue
        }

        l := i+1
        r := len(nums)-1

        for {
            
            if l >= r {
                break
            }

            three_sum := a + nums[l] + nums[r]
           
            // if the result is greater than 0, we need less
            // we can do that by shifting to left of sorted array
            if three_sum >0 {
                r -=1 
            } else if three_sum < 0 {
                // if result is less than 0, we need to increase value
                // we do that by moving to right of sorted array
                l+=1
            } else {
                // it means the value is equals 0
                res = append(res,[]int{a,nums[l],nums[r]})
                l +=1 
                for {
                    if (nums[l] != nums[l-1]) || (l >=r) {
                        break
                    }
                    l += 1
                }
            }
        }
    }
    return res
}
func main() {
  fmt.Println(threeSum([]int{-1,0,1,2,-1,-4} ))
}
