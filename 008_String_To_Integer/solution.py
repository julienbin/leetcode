class Solution(object):
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        str = str.strip()
        if not str:
            return 0
            
        MAX_INT = 2147483647
        MIN_INT = -2147483648
        ret = 0
        pos = 0
        overflow = False
        sign = 1
        
        if str[pos] == '-':
            sign = -1
            pos += 1
        elif str[pos] == '+':
            pos += 1
            
        for i in range(pos, len(str)):
            if not str[i].isdigit():
                break;
            ret = ret * 10 + int(str[i])
            if not MIN_INT <= sign * ret <= MAX_INT:
                overflow = True
                break
        if overflow:
            return MAX_INT if sign == 1 else MIN_INT
        else:
            return sign * ret
