class Solution:
    def uniqueOccurrences(self, arr: list[int]) -> bool:
        freq = {}

        for i in arr:
            if i in freq:
                freq[i] += 1
            else:
                freq[i] = 1

        oc = set()

        for k in freq.values():
            if k in oc:
                return False
            oc.add(k)
        return True


s = Solution()
print(s.uniqueOccurrences([1, 2, 2, 3, 3, 3]))
