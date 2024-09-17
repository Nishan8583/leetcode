/*
O(2n)
XOR = set to 1 if both different numbers
order does not matter, and number XORed with itself returns 0
Ex: 5^5 = 0, and 5^5^3 = 3, so if we get 2 arrays, and we have all the same number, we get 0, and if one is missing we get that number
[3,0,1] range [0,3]
(3^0^1) ^ (0,1,2,3) = 2

Another one,
sum the 2 array, and caclulate difference, and u get the value

So how was it different? how did he come  up with better solutions ?
1. He looked at the bigger impact, like took the whole array, and see what the difference would be. like sum
2. XOR is something i knew but i could not realize it lmao
*/
package main

import "fmt"

func missingNumber(nums []int) int {
	checked := make(map[int]bool, len(nums))
	for i := 0; i <= len(nums); i++ {
		checked[i] = false
	}

	for _, v := range nums {
		delete(checked, v) // previously i just set checked values to true, so loop below would take more time, this reduced speed alot
	}

	for i, v := range checked {
		if !v {
			return i
		}
	}

	return -1
}
func missingNumberXOR(nums []int) int {
	totalRange := make([]int, 0, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		totalRange = append(totalRange, i)
	}
	res := 0
	fmt.Println(totalRange)
	for _, v := range nums {
		res = res ^ v
	}
	fmt.Println(res)
	for _, v := range totalRange {
		res = res ^ v
	}

	return res
}

func main() {
	fmt.Println(missingNumberXOR([]int{3, 0, 1}))
}
