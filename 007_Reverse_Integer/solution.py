class Solution(object):
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        sign = 1 if x > 0 else -1
        x = abs(x)
        res = 0
        while x > 0:
            res = res * 10 + x % 10
            if res > math.pow(2, 31):
                return 0
            x = x / 10
        return res * sign
