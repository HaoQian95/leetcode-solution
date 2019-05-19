class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        start = 0
        l_len = 0
        res = 0
        ss = set([])
        for i in range(len(s)):
            if s[i] not in ss:
                ss.add(s[i])
                l_len += 1
            else:
                if l_len > res:
                    res = l_len
                while s[start]!=s[i]:
                    ss.remove(s[start])
                    start += 1
                start = start + 1
                l_len = i - start + 1
        return res if res > l_len else l_len