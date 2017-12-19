class Solution(object):
    def intToRoman(self, num):
        """
        :type num: int
        :rtype: str
        """
        syb = ['M', 'CM', 'D', 'CD', 
               'C', 'XC', 'L', 'XL',
               'X', 'IX', 'V', 'IV',
               'I']
        val = [1000, 900, 500, 400,
               100, 90, 50, 40,
               10, 9, 5, 4, 1]
        roma = ''
        i = 0
        while num > 0:
            for _ in range(num // val[i]):
                roma += syb[i]
                num -= val[i]
            i += 1
        return roma
