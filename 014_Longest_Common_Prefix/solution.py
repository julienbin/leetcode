class Solution(object):
    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        if strs == []:
            return ""
            
        if len(strs[0]) == 0:
            return ""
            
        i = -1
        while i < len(strs[0]):
            i += 1
            for str in strs:
                if len(str) <= i or str[i] != strs[0][i]:
                    return strs[0][:i]
        return ""
