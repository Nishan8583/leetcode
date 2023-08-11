/*
1011 %2, modulo works on the last binary digit it seems, so if it is one, it returns 1, else 0, so keep on doing it and shift bits to right
1011 % 2 = 1 ,so increase number of 1 bits then shift to right
101 %2 = 1, again increase and shift
10 %2 = 0 so no increase
1 %2 = 1 so increase and shift
check to 0 for the input is 0 so we exit
*/

func hammingWeight(num uint32) int {
    var res uint32 = 0
    for {
        if num == 0 {
            break
        }
        res  += num %2
        num = num >> 1
    }
    return int(res)
}
