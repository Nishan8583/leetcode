/*
I do not get good speed for some reason
visualize
[100,4,200,1,3,2]

x[1,2,3,4]		x[100]		x[200]
---------------------------------
1. Put all them numbers in hashmap
2. check if left neigbours exist, if so , its not the begining, so break out
3. check if right neigbour exist, if so increase number
*/

func longestConsecutive(nums []int) int {
	sets := map[int]struct{}{}
	for _, v := range nums {
		sets[v] = struct{}{}
	}

	max_c := 0
	for _, n := range nums {
		c := 1
		_, ok := sets[n-1]
		rn := n + 1
		if !ok {
			for {
				if _, ok := sets[rn]; ok {
					rn += 1
					c += 1
				} else {
					break
				}
			}
		}
		if c > max_c {
			max_c = c
		}
	}
	return max_c
}
