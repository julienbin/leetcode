class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        i = 0
        j = 0
        max = 0
        exist = {}
        for n in range(len(s)):
            exist[s[n]] = -1
        while j < len(s):
            if exist[s[j]] == -1:
                exist[s[j]] = 1
                j += 1
            else:
                while s[i] != s[j]:
                    exist[s[i]] = -1
                    i += 1
                i += 1
                j += 1
            max = max if max > (j - i) else (j - i)
        return max
