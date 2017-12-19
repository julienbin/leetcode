class Solution(object):
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        if x < 0:
            return False
        origin = x
        reverse = 0
        while x > 0:
            reverse = reverse * 10 + x % 10
            if reverse >= math.pow(2, 31):
                return False
            x = x / 10
        return (reverse == origin) 
