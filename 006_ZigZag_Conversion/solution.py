class Solution(object):
    def convert(self, s, numRows):
        """
        :type s: str
        :type numRows: int
        :rtype: str
        """
        if numRows == 1:
            return s
        tmp = ['' for i in range(numRows)]
        index = -1
        step = 1
        for i in range(len(s)):
            index += step
            if index == numRows:
                index -= 2
                step = -1
            elif index == -1:
                index = 1
                step = 1
            tmp[index] += s[i]
        return ''.join(tmp)
