class Solution:
    def removeStars(self, s: str) -> str:
        stack = []
        for c in s:
            if c == "*":
                stack.pop()
                continue
            else:
                stack.append(c)

        return "".join(stack)


s = Solution()
print(s.removeStars("leet**cod*e"))
