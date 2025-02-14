"""
I had misunderstood the problem it seems.
THe qestion is asking, which subscrtrnig in t can divide both s and t.
Solution:
1. Loop till the min of either s and t.
2. Check if till that position, both the strings can be generated if the substring is multipled by some value.
Ex: ABABAB and AB, RESULT AB can be multipled by 3 and 1 to produce both of them.
3. We determine this by calling divides for both str1 and str2.
4. Divides checks first the new substring range can divide the passed string i.e. str1 and str2
5. It then chekc
lets say we reach at AB,
then for ABABAB s = AB and  t= ABABAB
we check len(t) % len(s) == 0, here 6%2 == 0, so it is divisible,
then we check s * len(t)//len(s) = s * 6/2 = AB * 3 = ABABAB, thus the string is recreatble
"""


class Solution(object):
    def gcdOfStrings(self, str1, str2):
        """
        :type str1: str
        :type str2: str
        :rtype: str
        """

        def divides(s, t):
            """Check if string s divides string t."""
            return len(t) % len(s) == 0 and s * (len(t) // len(s)) == t

        result = ""
        # perhaps if we do it in reverse we will take less time ?
        for i in range(1, min(len(str1), len(str2)) + 1):
            candidate = str1[:i]
            if divides(candidate, str1) and divides(candidate, str2):
                result = candidate
        return result


solution = Solution()
print(solution.gcdOfStrings("ABCABC", "ABC"))  # Output: "ABC"
print(solution.gcdOfStrings("ABABAB", "ABAB"))  # Output: "AB"
print(solution.gcdOfStrings("LEET", "CODE"))  # Output: ""
