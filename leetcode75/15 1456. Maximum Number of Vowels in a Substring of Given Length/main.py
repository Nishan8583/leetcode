class Solution:
    def maxVowels(self, s: str, k: int) -> int:
        """
        # this is the unoptimized one, times out
        left = 0  # holds the left pointer
        right = k
        # prev_vowel_count = 0  # holds the vowel_count of previous window
        max_vowel_count = 0  # max vowel count
        vowel_set = ("a", "e", "i", "o", "u")
        prev_vowel_found = False
        prev_vowel_count = 0
        # we loop till right <= len(s) because range loop inside  will loop till right-1, if right = 3, it ill loop 0 , 1, 2 which is what we want
        while right <= len(s):
            vowel_count = 0
            for i in range(left, right):
                if s[i] in vowel_set:
                    vowel_count += 1
            if (
                vowel_count == k
            ):  # return, because vowel_count count can never be larger than the slide window
                return vowel_count
            if vowel_count > max_vowel_count:
                max_vowel_count = vowel_count
            left += 1
            right += 1
        return max_vowel_count
        """

        # in the below optimized Solution we do not count all the vowels in a window, instead when we are in a new window
        # we check if previous left value had a vowel, if so we decrease the vowel count, we need a variable to hold prevous value count for that
        left = 0  # holds the left pointer
        max_vowel_count = 0  # max vowel count
        vowel_set = ("a", "e", "i", "o", "u")

        # we loop till k because range will loop till k-1, if k = 3, it ill loop 0 , 1, 2 which is what we want
        vowel_count = 0
        for i in range(left, k):
            if s[i] in vowel_set:
                vowel_count += 1

        # if vowel_count == k, then return, because since window os of size k, no way vowel_count is going to be greater than k
        if vowel_count == k:
            return vowel_count

        prev_vowel_count = vowel_count  # update the prev_vowel_count
        max_vowel_count = vowel_count  # update max_vowel_count
        left = left + 1  # update the left pointer
        right = k  # update the right pointer

        vowel_count = 0  # reset vowel_count to 0
        while right < len(s):
            # if previously counted vowel is no longer in the window, decrease the count
            vowel_count = prev_vowel_count
            if s[left - 1] in vowel_set:
                vowel_count -= 1

            if s[right] in vowel_set:
                vowel_count += 1
            if vowel_count == k:
                return vowel_count

            if vowel_count > max_vowel_count:
                max_vowel_count = vowel_count
            prev_vowel_count = vowel_count
            left += 1
            right += 1
        return max_vowel_count


s = Solution()
print(s.maxVowels("abciiidef", 3))
print(s.maxVowels("aeiou", 2))
print(s.maxVowels("leetcode", 3))
