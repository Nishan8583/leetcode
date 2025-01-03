class Solution(object):
    def mergeAlternately(self, word1, word2):
        """
        :type word1: str
        :type word2: str
        :rtype: str
        """
        word1 =list(word1)
        word2 = list(word2)
        final = ""
        while len(word1) != 0 or len(word2) != 0:
            if len(word1) != 0:
                final = final + word1.pop(0)
            if len(word2) != 0:
                final = final+word2.pop(0)
        return final

s = Solution()
v = s.mergeAlternately("abc","pqr")
print(v)

