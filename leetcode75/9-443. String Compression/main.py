"""
kind of like 3 pointers
["a","b","b","b","b","b","b","b","b","b","b","b","b"]
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
