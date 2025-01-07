"""
1. Left Right pointer
2. Left and right pointer first points to 0 index of str1
3. Check both str1 and str2 R pointer
4. If they match, increase R pointer
5. Keep on doing this, till R is not equal
6. Update the result if the len is greater than before
7. Now increase L by 1. and reset R to the new L.
8. Repeat the process
"""


class Solution(object):
    def gcdOfStrings(self, str1, str2):
        """
        :type str1: str
        :type str2: str
        :rtype: str
        """

        left = 0
        right = 0

        result = ""
        while left < len(str1):
            prevRight = right
            localValue = ""
            for i in str1[left:]:
                if right >= len(str2):
                    break

                # print("I", i, i, str2[right], right)
                if i == str2[right]:
                    # print("Strings match", i)
                    right += 1
                    localValue = localValue + i
                    # print(localValue)
                    if 
                    product = len(str1) / len(localValue)
                    if product % 1 != 0:
                        continue
                    product = int(product)
                    if localValue * product == str1 and (len(localValue) > len(result)):
                        result = localValue
                else:
                    # print("Breaking at", localValue)
                    break
            # print("ONE END", result)
            left += 1
            right = left
        return result
        # print(result)


s = Solution()
print(s.gcdOfStrings("ABCABC", "ABC"))
print(s.gcdOfStrings("ABABAB", "ABAB"))
print(s.gcdOfStrings("LEET", "CODE"))
