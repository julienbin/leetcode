class Solution(object):
        
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        if not s or len(s) == 1:  
            return s  
        # record the result value  
        max_length = 1  
        start_index = 0  
        end_index = 0  
        for i in range(0, len(s)-1):  
            count = 1  
            # aba  
            if s[i] != s[i+1]:  
                while i-count >= 0 and i + count < len(s) and s[i-count] == s[i+count]:  
                    count += 1  
                if (count-1) * 2 + 1 > max_length:  
                    max_length = (count-1) * 2 + 1  
                    start_index = i - count + 1  
                    end_index = i + count - 1  
            # xaaaaax  
            else:  
                count_repeat = 1  
                count = 0  
                while i+1 < len(s) and s[i] == s[i+1]:  
                    i += 1  
                    count_repeat += 1  
                while i-count_repeat+1-count >= 0 and i + count < len(s) and s[i-count_repeat+1-count] == s[i+count]:  
                    count += 1  
                if (count-1) * 2 + count_repeat > max_length:  
                    max_length = (count-1) * 2 + count_repeat  
                    start_index = i - count -count_repeat + 2  
                    end_index = i + count - 1  
        return s[start_index:end_index + 1]
