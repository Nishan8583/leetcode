class Solution:
    def reverseVowels(self, s: str) -> str:
        vowel_set = ("a", "A", "e", "E", "i", "I", "o", "O", "u", "U")
        stack = []
        for c in s:
            if c in vowel_set:
                stack.append(c)

        result = ""
        for c in s:
            if c in vowel_set:
                result = result + stack.pop()
            else:
                result = result + c
        return result


s = Solution()
print(s.reverseVowels("IceCreAm"))
