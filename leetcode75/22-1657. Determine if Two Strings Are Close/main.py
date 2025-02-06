class Solution:
    def closeStrings(self, word1: str, word2: str) -> bool:
        c1 = self.count_char(word1)
        c2 = self.count_char(word2)

        n1 = self.get_freq_count(c1)
        n2 = self.get_freq_count(c2)
        s1 = set(word1)
        s2 = set(word2)

        # if the each character appear the same number of times, then operation 1 can be done
        # If the number of frequencies count are same and the charactrers are same, then operation 2
        # can be done
        return c1 == c2 or (n1 == n2 and s1 == s2)

    def count_char(self, word1: str) -> dict:
        count = {}
        for i in word1:
            if i in count:
                count[i] += 1
            else:
                count[i] = 1
        return count

    def get_freq_count(self, char_count: dict) -> dict:
        freq_count = {}
        for count in char_count.values():
            if count in freq_count:
                freq_count[count] += 1
            else:
                freq_count[count] = 1
        return freq_count


s = Solution()
print(s.closeStrings("abc", "bca"))
