# Group anagrams together
# anagrams will have same number of letters in the string
# any order
# DS: Map[[26]char]=[]
# we go through another loop and group them togeter
import collections

class Solution(object):
    def groupAnagrams(self, strs):
        """
        :type strs: List[str]
        :rtype: List[List[str]]
        """

        # holds map, key is array of all alphabet, 0 means not present, 1 means present
        
        ans = collections.defaultdict(list)
        
        for s in strs:
            temp = [0] * 26
            for c in s:
                temp[ord(c)-ord('a')]+=1 # increase the corresponding character count

            # append the string to the same character count
            ans[tuple(temp)].append(s)
            
        
        return ans.values()

def groupAnagrams(strs):
        """
        :type strs: List[str]
        :rtype: List[List[str]]
        """

        # holds map, key is array of all alphabet, 0 means not present, 1 means present
        
        ans = collections.defaultdict(list)
        
        for s in strs:
            temp = [0] * 26
            for c in s:
                temp[ord(c)-ord('a')]+=1 # increase the corresponding character count

            # append the string to the same character count
            ans[tuple(temp)].append(s)
            
        
        return ans.values()

print(groupAnagrams(["eat","tea","tan","ate","nat","bat"]))