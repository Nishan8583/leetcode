/*
inner loop same as hamming weight
1011 %2, modulo works on the last binary digit it seems, so if it is one, it returns 1, else 0, so keep on doing it and shift bits to right
1011 % 2 = 1 ,so increase number of 1 bits then shift to right
101 %2 = 1, again increase and shift
10 %2 = 0 so no increase
1 %2 = 1 so increase and shift
check to 0 for the input is 0 so we exit
just add the number of 1 bits to array.

*/
func countBits(n int) []int {
    res := make([]int,0,n+1)
    for i:=0; i<=n; i++ {
      temp := i  
      c := 0
      for {
        if i == 0 {
          res = append(res,c)
          break
        }
        c += i %2
        i  = i >> 1
      }
      i = temp // prevent infinite loop, since i will always be 0 because of inner for loop, and ++  will always result in 1
    }
    return res
}