"""
最大回文子串
🔗： https://leetcode-cn.com/problems/longest-palindromic-substring/
"""


class Solution(object):
    P = list()

    def longestPalindrome(self, s):
        length = len(s)
        res = ''
        self.P = [[None for x in range(length)] for _ in range(length)]
        for x in range(len(s)):
            y = length
            while y > x:
                flag, left, right = self.is_palindrome(x, y - 1, s)
                if flag:
                    res = res if len(res) > len(s[left:right +
                                                  1]) else s[left:right + 1]
                    return res
                y -= 1

    def is_palindrome(self, left, right, s):
        """判断是否是回文"""
        if right - left <= 1:
            if not self.P[left][right]:
                self.P[left][right] = s[left] == s[right]
            return self.P[left][right], left, right
        self.P[left][right] = all([
            self.is_palindrome(left + 1, right - 1, s)[0], s[left] == s[right]
        ])
        return self.P[left][right], left, right


class Solution_V1(object):
    def longestPalindrome(self, s):
        length = len(s)
        res = s[:1]
        for mid in range(length):
            left_edge, right_edge = mid - 1, mid + 1

            # 查找有没有相邻元素相同， 将边界向右推
            while True:
                try:
                    if s[mid] == s[right_edge]:
                        right_edge += 1
                        res = s[mid:right_edge] if len(
                            s[mid:right_edge]) > len(res) else res
                    else:
                        break
                except IndexError:
                    break

            while True:

                try:
                    left, right = s[left_edge], s[right_edge]
                    if left == right:
                        res = s[left_edge:right_edge +
                                1] if len(s[left_edge:right_edge +
                                            1]) > len(res) else res
                        left_edge, right_edge = left_edge - 1, right_edge + 1
                    else:
                        break
                except IndexError:
                    break
        return res


class Solution_V2(object):
    def longestPalindrome(self, s):
        pass


if __name__ == "__main__":
    print(Solution_V1().longestPalindrome("bananas"))
