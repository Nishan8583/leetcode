class Solution:
    def findDifference(self, nums1: list[int], nums2: list[int]) -> list[list[int]]:
        n1_hashmap = {}  # hashmap that has the values in nums1_array
        n2_hashmap = {}
        for i in nums1:
            n1_hashmap[i] = 1

        for i in nums2:
            n2_hashmap[i] = 1

        answer1 = []
        answer2 = []
        for i in n1_hashmap.keys():
            if i not in n2_hashmap:
                answer1.append(i)

        for j in n2_hashmap.keys():
            if j not in n1_hashmap:
                answer2.append(j)

        return [answer1, answer2]


s = Solution()
print(s.findDifference([1, 2, 3], [2, 4, 6]))
