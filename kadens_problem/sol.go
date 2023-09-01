package main

import (
	"log"
)

func kaden(nums []int) (int,int,int){
  if len(nums) == 0 {
    return 0,0,0
  }

  max_sum := nums[0]
  curr_sum := -11111111111
  max_L := 0
  max_R := 0

  L :=0 
  for R:=0; R<len(nums); R++ {

    if curr_sum < 0 {
      curr_sum = 0
      L = R
    }
    curr_sum += nums[R]

    if curr_sum > max_sum {
     max_sum = curr_sum
      max_L = L 
      max_R = R
    }
    
  }

  return max_sum, max_L,max_R
}



func main() {
	input := []int{4, -1, 2, -7 , 3, 4}
  if o,L,R := kaden(input);(o != 7 || L!=4 || R != 5) {
    log.Fatal("Unexpected output ",o, L, R)
  }
}
