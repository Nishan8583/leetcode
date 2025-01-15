"""
1. Basically 2 pointers
2. "i" points to first occurence of a character that has possibility of match, "j" points to last match.
3. We will keep increasing "j" till value char[i] and char[j] are same.
4. Once they are not same we will get the count i.e. j-i.
For example in ["a","b","b","b","b","b","b","b","b","b","b","b","b"], when i = 1, where b is the char, at end of inner
for loop j will point to final value. j-i will be 12.
5. "to" points to the place where we are going to update the string compression information.
For example when i = 0, j = 0, to will also be 0, then we will update the char[to] to i, increase "to"
Then for i = 1, and j will poitn at last, then we will update char[to], increase "to" then since match is greater than 1,
we convert match count to string, and for each element on we update char[to], and inrease to++
6.. We will update char to char[:to] and return the length.
["a","b","b","b","b","b","b","b","b","b","b","b","b"]

Extra thoughts:
initially i = 0, j = 0 and to  = 0,
i and j gives group, we keep on increasing j till j and i are not same,
then we get difference of j and i, that gives us count of occurence
when j = 1, value are not same, update chars[to] to te last value, i.e. chars[i]
update i to j
now i = 1 and j = 1, we keep on going till last place, then we get the number, and since
gt than 1, after updating chars[to], update the count with digit
update chars[:to]
return length

Reference: https://www.youtube.com/watch?v=IhJgguNiYYk
this was the best reference i could find

"""


class Solution:
    def compress(self, chars: list[str]) -> int:
        i = 0
        to = 0
        while i < len(chars):
            j = i
            while j < len(chars) and chars[j] == chars[i]:
                j += 1

            num = j - i
            chars[to] = chars[i]
            to += 1
            if num > 1:
                for digit in str(num):
                    chars[to] = digit
                    to += 1
            i = j
        chars = chars[:to]
        return to


s = Solution()
print(s.compress(["a", "a", "b", "b", "c", "c", "c"]))
print(s.compress(["a"]))
print(s.compress(["a", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b"]))
