class Solution:
    def reverseWords(self, s: str) -> str:
        cleaned = s.strip()
        l = cleaned.split(" ")
        result = ""
        for i in reversed(l):
            if len(i) == 0:
                continue
            print("appending {%s}" % i)
            result = result + i.strip() + " "

        return result.rstrip()


s = Solution()
print(s.reverseWords("the sky is blue"))
print(s.reverseWords("a good   example"))
