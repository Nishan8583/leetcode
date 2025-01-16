"""
1. Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
Ex: s = "abc", t = "ahbgdc", if len(t) < len(s), return false, its never going to work
2. left points to current char of s. in the beginning it is 0, right points to current char of t, in the beginning it is 0
3. loop through t, if t[right] == s[left], left +=1 and right +=1, else just increase right+=1
4. Check if left >= len(s) during the loop, then the values are already present, and just return fasle
5. Out of the loop, check left and if not return false
"""


class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        if len(s) == 0:
            return True
        left = 0
        for char in t:
            if s[left] == char:
                left += 1
            if left >= len(s):
                return True
        return left >= len(s)


s = Solution()
print(s.isSubsequence("abc", "ahbgdc"))
print(s.isSubsequence("axc", "ahbgdc"))
