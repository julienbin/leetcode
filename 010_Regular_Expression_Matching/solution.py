class Solution(object):
    def isMatch(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: bool
        """
        if not s:
            return not p or (len(p) > 1 and p[1] == '*' and self.isMatch(s,p[2:]))
        elif not p:
            return False
        
        if len(p) == 1:
            return len(s) == 1 and (p[0] == '.' or p[0] == s[0])
        
        if p[1] != '*':
            return (p[0] == '.' or p[0] == s[0]) and self.isMatch(s[1:], p[1:])
        else:
            if (s[0] != p[0] and p[0] != '.'):
                return self.isMatch(s, p[2:])
            else:
                if self.isMatch(s, p[2:]):
                    return True
                    
                i = 0;
                while i < len(s) and (s[i] == p[0] or p[0] == '.'):
                    if self.isMatch(s[i+1:], p[2:]):
                        return True
                    i += 1

                return False
