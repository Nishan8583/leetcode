package main

func search(nums []int, target int) int {
	l := 0             // leftmost index
	r := len(nums) - 1 // rightmost index

	for l <= r {
		mid := l + ((r - l) / 2) // mid index
		//fmt.Println(l,mid,r)
		if nums[mid] == target {
			return mid
		} else if target > nums[mid] {
			//fmt.Println("target is greater so shift l to right")
			l = mid + 1
			mid = l + ((r - l) / 2)
		} else if target < nums[mid] {
			//fmt.Println("shift r to left")
			r = mid - 1
			mid = l + ((r - l) / 2)
		}
	}

	return -1
}
