package main

import "fmt"

func pivotIndex(nums []int) int {
	totalSum := 0
	for _, v := range nums {
		totalSum += v
	}
	//fmt.Println("Total Sum", totalSum)
	prefixSum := 0
	for i, v := range nums {
		rightSum := totalSum - v - prefixSum
		if rightSum == prefixSum {
			return i
		}
		prefixSum += v
	}
	return -1

}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
}
