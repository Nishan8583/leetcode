"""
Gn the range [1, 300].

iven an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there will not be input like 3a or 2[4].

The test cases are generated so that the length of the output will never exceed 105.



Example 1:

Input: s = "3[a]2[bc]"
Output: "aaabcbc"

Example 2:

Input: s = "3[a2[c]]"
Output: "accaccacc"

Example 3:

Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"



Constraints:

    1 <= s.length <= 30
    s consists of lowercase English letters, digits, and square brackets '[]'.
    s is guaranteed to be a valid input.
    All the integers in s are i
"""

"""
Solution
3[a]2[bc]
1.Stack []
2.Check if it is number, if it is, a helper function
if i.isNum() retrieve_string(s[i+1:]), it gives, resulting string and it gives the index it finished at 
3

helper:
    1. check if first is [ 
    2. if it is, increase till we find ]
    3. return the string
    4. If within that loop, we find another number 
    then do a recursive call, return the string ,the string will be appended to the final

"""


class Solution:
    def decodeString(self, s: str) -> str:
        stack = []
        for i in range(0, len(s)):
            if s[i] != "]":
                stack.append(s[i])
            else:
                # do not append ]
                substr = ""
                while stack[-1] != "[":  # keep looping till we find [
                    substr = (
                        stack.pop() + substr
                    )  # since in reverse order append stack.pop first
                stack.pop()  # popping [

                k = ""
                while stack and stack[-1].isdigit():
                    k = stack.pop() + k

                stack.append(int(k) * substr)

        return "".join(stack)


s = Solution()
v1 = s.decodeString("3[a]2[bc]")
print(v1, v1 == "aaabcbc")
v2 = s.decodeString("3[a2[c]]")
print(v2, v2 == "accaccacc")
v3 = s.decodeString("2[abc]3[cd]ef")
print(v3, v3 == "abcabccdcdcdef")
