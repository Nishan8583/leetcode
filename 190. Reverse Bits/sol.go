/*
A bit confusing  explainnation to  be honest
right shift to the ith position to get  each  bits  
& with 1 so thjat we can check wether it is 0 or 1
shift bit to the reverse ith position, OR it to res to  put it there
*/

package main
func reverseBits(num uint32) uint32 {
  var res uint32  = 0
  for i:=0; i < 32;  i++ {
    bit  := (num >> i) & 1
    res = res | (bit << (31  - i))
  }
  return res
}
