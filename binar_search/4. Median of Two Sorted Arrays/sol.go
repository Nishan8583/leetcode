// leetcode: https://leetcode.com/problems/median-of-two-sorted-arrays/description/
// Did not need to check neetcode solution for this hard one suprisingly ?
// array are sorted, loop through both array, go by each element
// if current element in right array is > than current element in left array, append element from left array
// and increase l, cause current l was just appended
// else append r and increase r
// go through rest of element in arrays and append, only one of the array will have any values, so we are good
// if the mid value is even (i.e. len(new_arr) == 2 || (len(new_arr)%2 == 0)), we need to do (arr[mid]+arr[mid-1])/2 (i.e.e ),else return the midvalue
package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := 0
	r := 0

	new_arr := []int{}

	for {
		if l >= len(nums1) || r >= len(nums2) {
			//fmt.Println("breaking out in ", l, r)
			break
		}

		if nums2[r] > nums1[l] {
			new_arr = append(new_arr, nums1[l])
			l++
			continue
		} else {
			new_arr = append(new_arr, nums2[r])
			r++
			continue
		}
	}

	//fmt.Println(l)
	if l < (len(nums1)) {
		for i := l; i < len(nums1); i++ {
			//fmt.Println("appending for nums1", i)
			new_arr = append(new_arr, nums1[i])
		}
	}
	//fmt.Println(r)
	if r < (len(nums2)) {
		for i := r; i < len(nums2); i++ {
			//fmt.Println("appending for nums2", i)
			new_arr = append(new_arr, nums2[i])
		}
	}
	fmt.Println(new_arr)
	mid := len(new_arr) / 2

	return_value := 0.0
	if len(new_arr) < 2 {
		return float64(new_arr[0])
	}
	// if the mid value is even
	//fmt.Println("mid value is", mid)
	if len(new_arr) == 2 || (len(new_arr)%2 == 0) {
		//fmt.Println("Mid value is even, dividing the sum of ", mid, new_arr[mid], new_arr[mid-1])
		return_value = (float64(new_arr[mid]) + float64(new_arr[mid-1])) / 2
	} else {
		return_value = float64(new_arr[mid])
	}
	return return_value
}

func main() {
	fmt.Println("answer 1", findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println("answer 2", findMedianSortedArrays([]int{1, 3}, []int{2, 4}))
	fmt.Println("answer 3", findMedianSortedArrays([]int{1, 3, 6, 8}, []int{5, 6, 7, 8}))
	fmt.Println("answer 4", findMedianSortedArrays([]int{}, []int{1}))
	fmt.Println("answer 5", findMedianSortedArrays([]int{}, []int{2, 3}))
	fmt.Println("answer 6", findMedianSortedArrays([]int{}, []int{1, 2, 3, 4, 5}))
}
